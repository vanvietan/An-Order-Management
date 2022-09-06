package user

import (
	"context"
	"errors"
	"order-mg/db"
	"order-mg/internal/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteUser(t *testing.T) {
	type arg struct {
		givenID   int64
		expResult bool
		expErr    error
	}

	tcs := map[string]arg{
		"success: delete success": {
			givenID:   101,
			expResult: true,
		},
		"fail: no user id": {
			givenID:   200,
			expResult: true,
			expErr:    errors.New("record not found"),
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
			rs, err := instance.DeleteUser(context.Background(), tc.givenID)

			//THEN
			if err != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.Equal(t, tc.expResult, rs)
			}

		})
	}

}
