package user

import (
	"context"
	mocks "order-mg/internal/mocks/repository/user"
	"order-mg/internal/model"
	"order-mg/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
)

func TestUpdateUser(t *testing.T) {
	type updateUser struct {
		mockID    int64
		mockInput model.Users
		mockResp  model.Users
		mockErr   error
	}
	type arg struct {
		givenID    int64
		updateUser updateUser
		givenInput model.Users
		expRs      model.Users
		expErr     error
	}

	tcs := map[string]arg{
		"success: update with no error": {
			givenID: 0,
			updateUser: updateUser{
				mockID: 0,
				mockInput: model.Users{
					Id:          0,
					Name:        "nghia",
					Username:    "abc",
					Password:    "nghia",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "ADMIN",
					CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				},
				mockResp: model.Users{
					Id:          0,
					Name:        "nghia",
					Username:    "abc",
					Password:    "nghia",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "ADMIN",
					CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				},
			},
			givenInput: model.Users{
				Id:          0,
				Name:        "nghia",
				Username:    "abc",
				Password:    "nghia",
				PhoneNumber: "123",
				Address:     "abc",
				Age:         1,
				Role:        "ADMIN",
				CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
			expRs: model.Users{
				Id:          0,
				Name:        "nghia",
				Username:    "abc",
				PhoneNumber: "123",
				Address:     "abc",
				Age:         1,
				Role:        "ADMIN",
				CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
		},
	}

	ctx := context.Background()

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.UserRepository)
			instance.On("GetUserByID", mock.Anything, tc.updateUser.mockID).Return(tc.updateUser.mockResp, tc.updateUser.mockErr)
			instance.On("UpdateUser", mock.Anything, tc.updateUser.mockInput).Return(tc.updateUser.mockResp, tc.updateUser.mockErr)

			hashPasswordFunc = func(s string) string {
				return "nghia"
			}

			defer func() {
				hashPasswordFunc = util.HashPassword
			}()

			//WHEN
			svc := New(instance)
			rs, err := svc.UpdateUser(ctx, tc.givenInput, tc.givenID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRs, rs)
				// if !cmp.Equal(tc.expRs, rs,
				// 	cmpopts.IgnoreFields(model.Users{}, "Password")) {
				// 	t.Errorf("\n user mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expRs, rs,
				// 		cmp.Diff(tc.expRs, rs, cmpopts.IgnoreFields(model.Users{}, "Password")))
				// 	t.FailNow()
				// }
			}
		})
	}
}
