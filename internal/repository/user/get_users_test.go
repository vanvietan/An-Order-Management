package user

import (
	"context"
	"order-mg/db"
	"order-mg/internal/model"
	"order-mg/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	type arg struct {
		givenLimit  int
		givenLastID int64
		expResult   []model.Users
		expErr      error
	}

	tcs := map[string]arg{
		"success: gets all": {
			givenLimit:  3,
			givenLastID: 0,
			expResult: []model.Users{
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
				{
					Id:          99,
					Name:        "nghia",
					Username:    "nghia",
					Password:    "nghia",
					PhoneNumber: "123",
					Address:     "nghia",
					Age:         1,
					Role:        "USER",
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
		},
		"empty result": {
			givenLimit:  3,
			givenLastID: 1,
			expResult:   []model.Users{},
		},
		"success: last id": {
			givenLimit:  3,
			givenLastID: 101,
			expResult: []model.Users{
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
				{
					Id:          99,
					Name:        "nghia",
					Username:    "nghia",
					Password:    "nghia",
					PhoneNumber: "123",
					Address:     "nghia",
					Age:         1,
					Role:        "USER",
					CreatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	dbConn, errDB := db.CreateDBConnection()
	require.NoError(t, errDB)

	errExe := util.ExecuteTestData(dbConn, "./testdata/get_users.sql")
	require.NoError(t, errExe)

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			// GIVEN
			instance := New(dbConn)

			// WHEN
			rs, err := instance.GetUsers(context.Background(), tc.givenLimit, tc.givenLastID)

			// THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rs)
			}
		})
	}
}
