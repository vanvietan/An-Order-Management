package user

import (
	"context"
	"errors"
	mocks "order-mg/internal/mocks/repository/user"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
)

func TestDeleteUser(t *testing.T) {
	type deleteUser struct {
		mockID  int64
		mockErr error
	}
	type arg struct {
		givenID    int64
		deleteUser deleteUser
		expErr     error
	}

	tcs := map[string]arg{
		"fail: id isn't existed": {
			givenID: 200,
			deleteUser: deleteUser{
				mockID:  200,
				mockErr: errors.New("can't delete a user"),
			},
			expErr: errors.New("can't delete a user"),
		},
		"success: delete successful": {
			givenID: 101,
			deleteUser: deleteUser{
				mockID: 101,
			},
		},
	}

	ctx := context.Background()

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.UserRepository)
			instance.On("DeleteUser", mock.Anything, tc.deleteUser.mockID).Return(tc.deleteUser.mockErr)

			//WHEN
			svc := New(instance)
			err := svc.DeleteUser(ctx, tc.givenID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
