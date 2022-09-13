package test

import (
	"context"
	mocks "order-mg/internal/mocks/repository/user"
	"order-mg/internal/model"
	"testing"
	"time"
)

func TestGetUsers(t *testing.T) {
	type arg struct {
		givenLimit  int
		givenLastID int64
		mockResp    []model.Users
		mockErr     error
		expErr      error
	}

	tcs := map[string]arg{
		"success: get users": {
			givenLimit:  3,
			givenLastID: 0,
			mockResp: []model.Users{
				{
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
				{
					Id:          100,
					Name:        "abc",
					Username:    "abc",
					Password:    "abc",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "ADMIN",
					CreatedAt:   time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				},
			},
		},
		"empty result": {
			givenLimit:  3,
			givenLastID: 1,
			mockResp:    []model.Users{},
		},
		"success: last id": {
			givenLimit:  3,
			givenLastID: 200,
			mockResp: []model.Users{
				{
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
				{
					Id:          100,
					Name:        "abc",
					Username:    "abc",
					Password:    "abc",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "ADMIN",
					CreatedAt:   time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	ctx := context.Background()

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.UserRepository)
			instance.On("GetUsers", ctx, tc.givenLimit, tc.givenLastID)

			//WHEN

			//THEN
		})
	}
}
