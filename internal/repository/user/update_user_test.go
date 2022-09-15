package user

import (
	"context"
	"errors"
	"order-mg/db"
	"order-mg/internal/model"
	"order-mg/internal/util"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

func TestUpdateUser(t *testing.T) {
	type arg struct {
		givenResult model.Users
		expResult   model.Users
		expErr      error
	}
	tcs := map[string]arg{
		"success: update with no error": {
			givenResult: model.Users{
				Id:          100,
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
			expResult: model.Users{
				Id:          100,
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
		"fail: update fail": {
			givenResult: model.Users{
				Id:          103,
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
			expResult: model.Users{},
			expErr:    errors.New("ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)"),
		},
	}

	dbConn, errDB := db.CreateDBConnection()
	require.NoError(t, errDB)

	errExe := util.ExecuteTestData(dbConn, "./testdata/get_users.sql")
	require.NoError(t, errExe)

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := New(dbConn)

			//WHEN
			rs, err := instance.UpdateUser(context.Background(), tc.givenResult)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				// require.Equal(t, tc.expResult, rs)
				if !cmp.Equal(tc.expResult, rs,
					cmpopts.IgnoreFields(model.Users{}, "CreatedAt", "UpdatedAt")) {
					t.Errorf("\n user mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expResult, rs,
						cmp.Diff(tc.expResult, rs, cmpopts.IgnoreFields(model.Users{}, "CreatedAt", "UpdatedAt")))
					t.FailNow()
				}
			}
		})
	}

}
