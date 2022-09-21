package user

import (
	"context"
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
		NOBODY

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
		givenLimit         string
		givenCursor        string
		getUsers           getUsers
		getUsersMockCalled bool
		expRs              string
		expHTTPCode        int
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
			getUsersMockCalled: true,
			expHTTPCode:        http.StatusOK,
		},
		"fail: invalid limit": {
			givenLimit:         "abc",
			givenCursor:        "0",
			getUsersMockCalled: false,
			expRs:              `{"code":"invalid_request", "description":"limit must be a number"}`,
			expHTTPCode:        http.StatusBadRequest,
		},
		"fail: invalid lastID": {
			givenLimit:         "20",
			givenCursor:        "abc",
			getUsersMockCalled: false,
			expRs:              `{"code":"invalid_request", "description":"cursor must be a number"}`,
			expHTTPCode:        http.StatusBadRequest,
		},
		"success: empty ": {
			givenLimit:         "3",
			givenCursor:        "1",
			getUsersMockCalled: true,
			getUsers: getUsers{
				mockLimit:  3,
				mockCursor: 1,
				mockOut:    []model.Users{},
				mockErr:    nil,
			},
			expRs:       `{"users":null,"cursor":0}`,
			expHTTPCode: http.StatusOK,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.UserService)
			if tc.getUsersMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("GetUsers", mock.Anything, tc.getUsers.mockLimit, tc.getUsers.mockCursor).
						Return(tc.getUsers.mockOut, tc.getUsers.mockErr),
				}
			}
			//GIVEN
			path := "/users" + "?limit=" + tc.givenLimit + "&cursor=" + tc.givenCursor
			req := httptest.NewRequest(http.MethodGet, path, nil)
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.GetUsers(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
