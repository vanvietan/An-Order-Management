package user

import (
	"context"
	"errors"
	mocks "order-mg/internal/mocks/repository/user"
	"order-mg/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	type getUserByID struct {
		mockIn   int64
		mockResp model.Users
		mockErr  error
	}

	type arg struct {
		givenID     int64
		getUserByID getUserByID
		expRs       model.Users
		expErr      error
	}

	tcs := map[string]arg{
		"success: get a user": {
			givenID: 101,
			getUserByID: getUserByID{
				mockIn: 101,
				mockResp: model.Users{
					Id:          101,
					Name:        "abc",
					Username:    "abc1",
					Password:    "abc",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "ADMIN",
					CreatedAt:   time.Date(2022, 3, 15, 16, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 15, 16, 0, 0, 0, time.UTC),
				},
			},
			expRs: model.Users{
				Id:          101,
				Name:        "abc",
				Username:    "abc1",
				PhoneNumber: "123",
				Address:     "abc",
				Age:         1,
				Role:        "ADMIN",
				CreatedAt:   time.Date(2022, 3, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 15, 16, 0, 0, 0, time.UTC),
			},
		},
		"fail: invalid id": {
			givenID: 0,
			expRs:   model.Users{},
			expErr:  errors.New("invalid userID"),
		},
		"fail: can't find the id": {
			givenID: 200,
			getUserByID: getUserByID{
				mockIn:   200,
				mockResp: model.Users{},
				mockErr:  errors.New("invalid userID"),
			},
			expRs:  model.Users{},
			expErr: errors.New("invalid userID"),
		},
	}
	ctx := context.Background()

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.UserRepository)
			instance.On("GetUserByID", ctx, tc.getUserByID.mockIn).Return(tc.getUserByID.mockResp, tc.getUserByID.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.GetUserByID(ctx, tc.givenID)

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
