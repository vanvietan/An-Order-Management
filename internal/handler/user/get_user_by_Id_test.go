package user

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	mocks "order-mg/internal/mocks/service/user"
	"order-mg/internal/model"
	"testing"
	"time"
)

func TestGetUserByID(t *testing.T) {
	/*
		GET /users/{id}
		NO BODY

		need: id

		expect:
			+ http status code
			+ resp

		mock service

	*/

	type getUserByIDMock struct {
		mockIn  int64
		mockOut model.Users
		mockErr error
	}

	type arg struct {
		givenID               string
		getUserByIDMock       getUserByIDMock
		getUserByIDMockCalled bool
		expRs                 string
		expHTTPCode           int
	}

	tcs := map[string]arg{
		"success": {
			givenID: "12234",
			getUserByIDMock: getUserByIDMock{
				mockIn: 12234,
				mockOut: model.Users{
					Id:          12234,
					Name:        "abc",
					Username:    "abc",
					PhoneNumber: "0123456789",
					Address:     "tphcm",
					Age:         12,
					Role:        model.RoleAdmin,
					CreatedAt:   time.Date(2022, 9, 20, 11, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 9, 20, 11, 0, 0, 0, time.UTC),
				},
			},
			expRs: `{
			  "id": 12234,
			  "name": "abc",
			  "username": "abc",
			  "phone_number": "0123456789",
			  "address": "tphcm",
			  "age": 12,
			  "role": "ADMIN",
			  "created_at": "2022-09-20T11:00:00Z",
			  "updated_at": "2022-09-20T11:00:00Z"
			}`,
			getUserByIDMockCalled: true,
			expHTTPCode:           http.StatusOK,
		},
		"fail: invalid ID": {
			givenID:               "abcd",
			getUserByIDMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"id must be a number"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: can't find userID": {
			givenID:               "103",
			getUserByIDMockCalled: true,
			getUserByIDMock: getUserByIDMock{
				mockIn:  103,
				mockOut: model.Users{},
				mockErr: errors.New("something wrong"),
			},
			expRs:       `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// Mock
			mockSvc := new(mocks.UserService)
			if tc.getUserByIDMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("GetUserByID", mock.Anything, tc.getUserByIDMock.mockIn).
						Return(tc.getUserByIDMock.mockOut, tc.getUserByIDMock.mockErr),
				}
			}
			// Given
			req := httptest.NewRequest(http.MethodGet, "/users/", nil)
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			// When
			instance := New(mockSvc)
			instance.GetUserByID(res, req)
			// Then
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}

}
