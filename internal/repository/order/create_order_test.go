package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"order-mg/db"
	"order-mg/internal/model"
	"order-mg/internal/util"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	type arg struct {
		givenInput model.Order
		expResult  model.Order
		expErr     error
	}

	tcs := map[string]arg{
		"success: ": {
			givenInput: model.Order{
				Id:            103,
				Name:          "abc",
				Description:   "abc",
				TotalPrice:    100,
				Quantity:      100,
				Discount:      10,
				Shipping:      "abc",
				Status:        model.StatusApproved,
				UserId:        100,
				DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
				CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
			},
			expResult: model.Order{
				Id:            103,
				Name:          "abc",
				Description:   "abc",
				TotalPrice:    100,
				Quantity:      100,
				Discount:      10,
				Shipping:      "abc",
				Status:        model.StatusApproved,
				UserId:        100,
				DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
				CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
			},
		},
		"fail: duplicated": {
			givenInput: model.Order{
				Id:            100,
				Name:          "abc",
				Description:   "abc",
				TotalPrice:    100,
				Quantity:      100,
				Discount:      10,
				Shipping:      "abc",
				Status:        model.StatusApproved,
				UserId:        100,
				DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
				CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
			},
			expResult: model.Order{},
			expErr:    errors.New("ERROR: duplicate key value violates unique constraint \"orders_pkey\" (SQLSTATE 23505)"),
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
			rs, err := instance.CreateOrder(context.Background(), tc.givenInput)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rs)
			}
		})
	}
}
