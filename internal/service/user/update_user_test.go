package user

import (
	"context"
	mocks "order-mg/internal/mocks/repository/user"
	"order-mg/internal/model"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

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
					Name: "nghia",
					// Username:    "abc",
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
				Name:        "nghia",
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
			instance.On("GetUserByID", ctx, tc.updateUser.mockID).Return(model.Users{}, tc.updateUser.mockErr)
			instance.On("UpdateUser", ctx, tc.updateUser.mockInput).Return(tc.updateUser.mockResp, tc.updateUser.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.UpdateUser(ctx, tc.givenInput, tc.givenID)

			//THEN
			if err != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				// require.Equal(t, tc.expRs, rs)
				if !cmp.Equal(tc.expRs, rs,
					cmpopts.IgnoreFields(model.Users{}, "CreatedAt", "UpdatedAt", "Password")) {
					t.Errorf("\n user mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expRs, rs,
						cmp.Diff(tc.expRs, rs, cmpopts.IgnoreFields(model.Users{}, "CreatedAt", "UpdatedAt", "Password")))
					t.FailNow()
				}
			}
		})
	}
}
