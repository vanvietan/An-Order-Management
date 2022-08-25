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

func TestGetUserByID(t *testing.T) {
	type arg struct {
		givenId   int64
		expResult model.Users
		expErr    error
	}

	tcs := map[string]arg{
		"success: given id isn't 0": {
			givenId: 101,
			expResult: model.Users{
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
	}

	dbConn, errDb := db.CreateDBConnection()
	require.NoError(t, errDb)

	errExe := util.ExecuteTestData(dbConn, "./testdata/get_users.sql")
	require.NoError(t, errExe)

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GiVEN
			instance := New(dbConn)

			//WHEN
			rs, err := instance.GetUserByID(context.Background(), tc.givenId)

			//THEN
			if err != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.Equal(t, tc.expResult, rs)
			}
		})
	}
}
