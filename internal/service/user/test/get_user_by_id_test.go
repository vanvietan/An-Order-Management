package test

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
	type arg struct {
		givenID  int64
		mockErr  error
		mockResp model.Users
		expErr   error
	}

	tcs := map[string]arg{
		"success: get a user": {
			givenID: 101,
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
		"fail: id isn't existed": {
			givenID:  0,
			mockResp: model.Users{},
			expErr:   errors.New("invalid userID"),
		},
	}
	ctx := context.Background()

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.UserRepository)
			instance.On("GetUserByID", ctx, tc.givenID).Return(tc.mockResp, tc.mockErr)

			//WHEN
			rs, err := instance.GetUserByID(context.Background(), tc.givenID)

			//THEN
			if err != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.Equal(t, tc.mockResp, rs)
			}
		})
	}
}
