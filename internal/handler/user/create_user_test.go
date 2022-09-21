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
	"strings"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {

	type createUser struct {
		mockIn  model.Users
		mockOut model.Users
		mockErr error
	}
	type arg struct {
		givenBody            string
		createUser           createUser
		createUserMockCalled bool
		expRs                string
		expHTTPCode          int
	}
	tcs := map[string]arg{
		"success: ": {
			givenBody: `{
				"name": "abc",
				"username": "abc",
				"password": "abc",
				"phone_number": "0906312911",
				"address": "abc",
				"age": 90,
				"role": "USER"
				}`,
			createUser: createUser{
				mockIn: model.Users{
					Name:        "abc",
					Username:    "abc",
					Password:    "abc",
					PhoneNumber: "0906312911",
					Address:     "abc",
					Age:         90,
					Role:        "USER",
				},
				mockOut: model.Users{
					Id:          100,
					Name:        "abc",
					Username:    "abc",
					Password:    "abc",
					PhoneNumber: "0906312911",
					Address:     "abc",
					Age:         90,
					Role:        "USER",
					CreatedAt:   time.Date(2022, 9, 21, 16, 30, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 9, 21, 16, 30, 0, 0, time.UTC),
				},
			},
			createUserMockCalled: true,
			expRs: `{
				"id": 100,
				"name": "abc",
				"username": "abc",
				"phone_number": "0906312911",
				"address": "abc",
				"age": 90,
				"role": "USER",
				"created_at": "2022-09-21T16:30:00Z",
				"updated_at": "2022-09-21T16:30:00Z"
				}`,
			expHTTPCode: http.StatusOK,
		},
		"fail:invalid Body user name invalid": {
			givenBody: `{
				"name": "",
				"username": "abc",
				"password": "abc",
				"phone_number": "0906312911",
				"address": "abc",
				"age": 90,
				"role": "USER"
				}`,
			createUserMockCalled: false,
			expRs:                `{"code":"invalid request", "description":"invalid name"}`,
			expHTTPCode:          http.StatusBadRequest,
		},
		"fail:invalid Body user username invalid": {
			givenBody: `{
				"name": "abc",
				"username": "",
				"password": "abc",
				"phone_number": "0906312911",
				"address": "abc",
				"age": 90,
				"role": "USER"
				}`,
			createUserMockCalled: false,
			expRs:                `{"code":"invalid request", "description":"invalid username"}`,
			expHTTPCode:          http.StatusBadRequest,
		},
		"fail:invalid Body user password invalid": {
			givenBody: `{
				"name": "abc",
				"username": "abc",
				"password": "",
				"phone_number": "0906312911",
				"address": "abc",
				"age": 90,
				"role": "USER"
				}`,
			createUserMockCalled: false,
			expRs:                `{"code":"invalid request", "description":"password is invalid"}`,
			expHTTPCode:          http.StatusBadRequest,
		},
		"fail:invalid Body user address invalid": {
			givenBody: `{
				"name": "abc",
				"username": "abc",
				"password": "abc",
				"phone_number": "0906312911",
				"address": "",
				"age": 90,
				"role": "USER"
				}`,
			createUserMockCalled: false,
			expRs:                `{"code":"invalid request", "description":"address is invalid"}`,
			expHTTPCode:          http.StatusBadRequest,
		},

		"fail:invalid Body user age invalid": {
			givenBody: `{
				"name": "abc",
				"username": "abc",
				"password": "abc",
				"phone_number": "0906312911",
				"address": "abc",
				"age": 200,
				"role": "USER"
				}`,
			createUserMockCalled: false,
			expRs:                `{"code":"invalid request", "description":"user age is invalid"}`,
			expHTTPCode:          http.StatusBadRequest,
		},
		"fail:invalid Body user phone number invalid": {
			givenBody: `{
				"name": "abc",
				"username": "abc",
				"password": "abc",
				"phone_number": "xyz",
				"address": "abc",
				"age": 90,
				"role": "USER"
				}`,
			createUserMockCalled: false,
			expRs:                `{"code":"invalid request", "description":"phone number is invalid"}`,
			expHTTPCode:          http.StatusBadRequest,
		},
		"fail:invalid Body user role invalid": {
			givenBody: `{
				"name": "abc",
				"username": "abc",
				"password": "abc",
				"phone_number": "0906312911",
				"address": "abc",
				"age": 90,
				"role": "PRESIDENT"
				}`,
			createUserMockCalled: false,
			expRs:                `{"code":"invalid request", "description":"user role is invalid"}`,
			expHTTPCode:          http.StatusBadRequest,
		},
		"fail: error from service": {
			givenBody: `{
				"name": "abc",
				"username": "abc",
				"password": "abc",
				"phone_number": "0906312911",
				"address": "abc",
				"age": 90,
				"role": "USER"
				}`,
			createUserMockCalled: true,
			createUser: createUser{
				mockIn: model.Users{
					Name:        "abc",
					Username:    "abc",
					Password:    "abc",
					PhoneNumber: "0906312911",
					Address:     "abc",
					Age:         90,
					Role:        "USER",
				},
				mockOut: model.Users{},
				mockErr: errors.New("something wrong"),
			},
			expRs:       `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.UserService)
			if tc.createUserMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("CreateUser", mock.Anything, tc.createUser.mockIn).
						Return(tc.createUser.mockOut, tc.createUser.mockErr),
				}
			}
			//GIVEN
			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.CreateUser(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
