package user

import (
	"context"
	"errors"
	mocks "order-mg/internal/mocks/repository/user"
	"order-mg/internal/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteUser(t *testing.T) {
	type deleteUser struct {
		mockIn     int64
		mockResult bool
		mockErr    error
	}
	type arg struct {
		givenID    int64
		deleteUser deleteUser
		expResult  bool
		expErr     error
	}

	tcs := map[string]arg{
		"fail: userID invalid": {
			givenID: 0,
			deleteUser: deleteUser{
				mockIn:     0,
				mockResult: false,
				mockErr:    errors.New("invalid userID"),
			},
			expResult: false,
			expErr:    errors.New("invalid userID"),
		},
		"fail: id isn't existed": {
			givenID: 200,
			deleteUser: deleteUser{
				mockIn:     200,
				mockResult: false,
				mockErr:    errors.New("userID is not existed"),
			},
			expResult: true,
			expErr:    errors.New("userID is not existed"),
		},
		"success: delete success": {
			givenID: 101,
			deleteUser: deleteUser{
				mockIn:     101,
				mockResult: true,
			},
			expResult: true,
		},
	}

	ctx := context.Background()

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.UserRepository)
			instance.On("GetUserByID", ctx, tc.deleteUser.mockIn).Return(model.Users{}, tc.deleteUser.mockErr)
			instance.On("DeleteUser", ctx, tc.deleteUser.mockIn).Return(tc.deleteUser.mockResult, tc.deleteUser.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.DeleteUser(ctx, tc.givenID)

			//THEN
			if err != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.Equal(t, tc.expResult, rs)
			}
		})
	}
}
