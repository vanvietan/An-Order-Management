package order

import (
	"context"
	"errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"order-mg/db"
	"order-mg/internal/model"
	"order-mg/internal/util"
	"testing"
	"time"
)

func TestUpdateOrder(t *testing.T) {
	type arg struct {
		givenInput model.Order
		expResult  model.Order
		expErr     error
	}

	tcs := map[string]arg{
		"success: ": {
			givenInput: model.Order{
				Id:            100,
				Name:          "ant",
				Description:   "ant",
				TotalPrice:    100,
				Quantity:      100,
				Discount:      10,
				Shipping:      "ant",
				Status:        model.StatusApproved,
				UserId:        100,
				DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
				CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
			},
			expResult: model.Order{
				Id:            100,
				Name:          "ant",
				Description:   "ant",
				TotalPrice:    100,
				Quantity:      100,
				Discount:      10,
				Shipping:      "ant",
				Status:        model.StatusApproved,
				UserId:        100,
				DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
				CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
			},
		},
		"fail: error in db wrong userID": {
			givenInput: model.Order{
				Id:            100,
				Name:          "ant",
				Description:   "ant",
				TotalPrice:    100,
				Quantity:      100,
				Discount:      10,
				Shipping:      "ant",
				Status:        model.StatusApproved,
				UserId:        0,
				DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
				CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
			},
			expResult: model.Order{},
			expErr:    errors.New("ERROR: insert or update on table \"orders\" violates foreign key constraint \"orders_user_id\" (SQLSTATE 23503)"),
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
			rs, err := instance.UpdateOrder(context.Background(), tc.givenInput)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				if !cmp.Equal(tc.expResult, rs,
					cmpopts.IgnoreFields(model.Order{}, "CreatedAt", "UpdatedAt")) {
					t.Errorf("\n order mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expResult, rs,
						cmp.Diff(tc.expResult, rs, cmpopts.IgnoreFields(model.Order{}, "CreatedAt", "UpdatedAt")))
					t.FailNow()
				}
			}
		})
	}
}
