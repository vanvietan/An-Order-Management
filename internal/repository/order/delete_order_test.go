package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"order-mg/db"
	"order-mg/internal/util"
	"testing"
)

func TestDeleteOrder(t *testing.T) {
	type arg struct {
		givenID int64
		expErr  error
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: 101,
		},
		"fail: ": {
			givenID: 200,
			expErr:  errors.New("record not found"),
		},
	}

	dbConn, errDB := db.CreateDBConnection()
	require.NoError(t, errDB)

	errExe := util.ExecuteTestData(dbConn, "./testdata/get_orders.sql")
	require.NoError(t, errExe)

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := New(dbConn)

			//WHEN
			err := instance.DeleteOrder(context.Background(), tc.givenID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}

		})
	}
}
