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
	"testing"
)

func TestDeleteUser(t *testing.T) {
	/*
		DELETE /users/{id}
		NO BODY
		need: id
		expect:
			+ http status code
			+ resp
		mock service
	*/

	type deleteUser struct {
		mockIn  int64
		mockErr error
	}
	type arg struct {
		givenID              string
		deleteUser           deleteUser
		deleteUserMockCalled bool
		expRs                string
		expHTTPCode          int
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: "123",
			deleteUser: deleteUser{
				mockIn:  123,
				mockErr: nil,
			},
			deleteUserMockCalled: true,
			expHTTPCode:          http.StatusOK,
			expRs:                `{"message":"Deleted User"}`,
		},
		"fail: invalid ID": {
			givenID: "abc",
			deleteUser: deleteUser{
				mockErr: errors.New("something wrong"),
			},
			deleteUserMockCalled: false,
			expHTTPCode:          http.StatusBadRequest,
			expRs:                `{"code":"invalid_request", "description":"id must be a number"}`,
		},
		"fail: cannot find a user to delete": {
			givenID: "123",
			deleteUser: deleteUser{
				mockIn:  123,
				mockErr: errors.New("something wrong"),
			},
			deleteUserMockCalled: true,
			expHTTPCode:          http.StatusInternalServerError,
			expRs:                `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.UserService)
			if tc.deleteUserMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("DeleteUser", mock.Anything, tc.deleteUser.mockIn).
						Return(tc.deleteUser.mockErr),
				}
			}

			//GIVEN
			req := httptest.NewRequest(http.MethodDelete, "/users/", nil)
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.DeleteUser(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
