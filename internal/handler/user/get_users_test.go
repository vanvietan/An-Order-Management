package user

import (
	"context"
	"fmt"
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

func TestGetUsers(t *testing.T) {
	/*
		GET /users
		NO BODY

		need: cursor , limit

		expect:
			+http status code
			+resp
		mock service
	*/

	type getUsers struct {
		mockLimit  int
		mockCursor int64
		mockOut    []model.Users
		mockErr    error
	}
	type arg struct {
		givenLimit            string
		givenCursor           string
		getUsers              getUsers
		getUserByIDMockCalled bool
		expRs                 string
		expHTTPCode           int
	}

	tcs := map[string]arg{
		"success:": {
			givenLimit:  "3",
			givenCursor: "0",
			getUsers: getUsers{
				mockLimit:  3,
				mockCursor: 0,
				mockOut: []model.Users{
					{
						Id:          101,
						Name:        "abc",
						Username:    "abc1",
						PhoneNumber: "123",
						Address:     "abc",
						Age:         1,
						Role:        "ADMIN",
						CreatedAt:   time.Date(2022, 9, 20, 11, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, 9, 20, 11, 0, 0, 0, time.UTC),
					},
					{
						Id:          100,
						Name:        "abc",
						Username:    "abc",
						PhoneNumber: "123",
						Address:     "abc",
						Age:         1,
						Role:        "ADMIN",
						CreatedAt:   time.Date(2022, 9, 20, 11, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, 9, 20, 11, 0, 0, 0, time.UTC),
					},
				},
			},
			expRs: `{
					"users": [
						{
							"id": 101,
							  "name": "abc",
							  "username": "abc1",
							  "phone_number": "123",
							  "address": "abc",
							  "age": 1,
							  "role": "ADMIN",
							  "created_at": "2022-09-20T11:00:00Z",
							  "updated_at": "2022-09-20T11:00:00Z"
						},
						{
							"id": 100,
							  "name": "abc",
							  "username": "abc",
							  "phone_number": "123",
							  "address": "abc",
							  "age": 1,
							  "role": "ADMIN",
							  "created_at": "2022-09-20T11:00:00Z",
							  "updated_at": "2022-09-20T11:00:00Z"
						}	
					],
					"cursor": 100
						}`,
			getUserByIDMockCalled: true,
			expHTTPCode:           http.StatusOK,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.UserService)
			if tc.getUserByIDMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("GetUsers", mock.Anything, tc.getUsers.mockLimit, tc.getUsers.mockCursor).
						Return(tc.getUsers.mockOut, tc.getUsers.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			routeCtx := chi.NewRouteContext()
			req.URL.Query().Add("limit", tc.givenLimit)
			req.URL.Query().Add("cursor", tc.givenCursor)
			
			fmt.Println(req.Context().Value("limit"))

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.GetUsers(res, req)

			//THEN
			//require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
