package user

import (
	"context"
	"errors"
	mocks "order-mg/internal/mocks/repository/user"
	"order-mg/internal/model"
	"order-mg/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	type createUser struct {
		mockInput model.Users
		mockResp  model.Users
		mockErr   error
	}
	type arg struct {
		createUser createUser
		givenInput model.Users
		expRs      model.Users
		expErr     error
	}

	tcs := map[string]arg{
		"success: create a user ": {
			createUser: createUser{
				mockInput: model.Users{
					Id:          103,
					Name:        "an",
					Username:    "an",
					Password:    "abc",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "USER",
					CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				},
				mockResp: model.Users{
					Id:          103,
					Name:        "an",
					Username:    "an",
					Password:    "abc",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "USER",
					CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				},
			},
			givenInput: model.Users{
				Name:        "an",
				Username:    "an",
				PhoneNumber: "123",
				Address:     "abc",
				Age:         1,
				Role:        "USER",
				CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
			expRs: model.Users{
				Id:          103,
				Name:        "an",
				Username:    "an",
				PhoneNumber: "123",
				Address:     "abc",
				Age:         1,
				Role:        "USER",
				CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
		},
		"fail: can't create a user": {
			createUser: createUser{
				mockInput: model.Users{
					Id:          103,
					Name:        "an",
					Username:    "an",
					Password:    "abc",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "USER",
					CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				},
				mockResp: model.Users{},
				mockErr:  errors.New("something error"),
			},
			givenInput: model.Users{
				Name:        "an",
				Username:    "an",
				PhoneNumber: "123",
				Address:     "abc",
				Age:         1,
				Role:        "USER",
				CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
			},
			expRs:  model.Users{},
			expErr: errors.New("something error"),
		},
		"fail: generate fail id": {
			createUser: createUser{
				mockInput: model.Users{
					Name:        "an",
					Username:    "an",
					Password:    "abc",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "USER",
					CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
				},
				mockResp: model.Users{},
				mockErr:  errors.New("something error"),
			},
			expRs:  model.Users{},
			expErr: errors.New("something error"),
		},
	}

	ctx := context.Background()

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.UserRepository)
			instance.On("CreateUser", mock.Anything, tc.createUser.mockInput).Return(tc.createUser.mockResp, tc.createUser.mockErr)

			getNextIDFunc = func() (int64, error) {
				if s == "fail: generate fail id" {
					return 0, errors.New("something error")
				}
				return 103, nil
			}
			hashPasswordFunc = func(s string) string {
				return "abc"
			}
			defer func() {
				getNextIDFunc = util.GetNextId
				hashPasswordFunc = util.HashPassword
			}()

			//WHEN
			svc := New(instance)
			rs, err := svc.CreateUser(ctx, tc.givenInput)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRs, rs)
			}
		})
	}
}
