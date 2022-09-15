package user

import (
	"context"
	"errors"
	mocks "order-mg/internal/mocks/repository/user"
	"order-mg/internal/model"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	type getUsers struct {
		mockLimit int
		mockIn    int64
		mockResp  []model.Users
		mockErr   error
	}
	type arg struct {
		givenLimit  int
		givenLastID int64
		getUsers    getUsers
		expRs       []model.Users
		expErr      error
	}

	tcs := map[string]arg{
		"success: get users": {
			givenLimit:  3,
			givenLastID: 0,
			getUsers: getUsers{
				mockLimit: 3,
				mockIn:    0,
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
			expRs: []model.Users{
				{
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
				{
					Id:          100,
					Name:        "abc",
					Username:    "abc",
					PhoneNumber: "123",
					Address:     "abc",
					Age:         1,
					Role:        "ADMIN",
					CreatedAt:   time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				},
			},
		},
		"fail: empty result": {
			givenLimit:  3,
			givenLastID: 1,
			getUsers: getUsers{
				mockLimit: 3,
				mockIn:    0,
				mockResp:  []model.Users{},
				mockErr:   errors.New("something wrong"),
			},
			expRs:  []model.Users{},
			expErr: errors.New("something wrong"),
		},
		"success: last id": {
			givenLimit:  3,
			givenLastID: 200,
			getUsers: getUsers{
				mockLimit: 3,
				mockIn:    0,
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
			expRs: []model.Users{
				{
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
				{
					Id:          100,
					Name:        "abc",
					Username:    "abc",
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
			instance.On("GetUsers", mock.Anything, tc.givenLimit, tc.givenLastID).Return(tc.getUsers.mockResp, tc.getUsers.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.GetUsers(ctx, tc.givenLimit, tc.givenLastID)

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
