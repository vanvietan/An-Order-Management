package order

import (
	"context"
	"github.com/stretchr/testify/require"
	"order-mg/db"
	"order-mg/internal/model"
	"order-mg/internal/util"
	"testing"
	"time"
)

func TestGetOrders(t *testing.T) {
	type arg struct {
		givenLimit int
		givenID    int64
		expResult  []model.Order
		expErr     error
	}
	tcs := map[string]arg{
		"success: get all orders": {
			givenLimit: 20,
			givenID:    0,
			expResult: []model.Order{
				{
					Id:            101,
					Name:          "abc1",
					Description:   "abc1",
					TotalPrice:    101,
					Quantity:      101,
					Discount:      11,
					Shipping:      "abc",
					Status:        model.StatusApproved,
					UserId:        101,
					DatePurchased: time.Date(2022, 3, 16, 0, 0, 0, 0, time.UTC),
					CreatedAt:     time.Date(2022, 3, 16, 16, 0, 0, 0, time.UTC),
					UpdatedAt:     time.Date(2022, 3, 16, 16, 0, 0, 0, time.UTC),
				},
				{
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
				{
					Id:            99,
					Name:          "an",
					Description:   "an",
					TotalPrice:    99,
					Quantity:      99,
					Discount:      9,
					Shipping:      "hcm",
					Status:        model.StatusShipped,
					UserId:        99,
					DatePurchased: time.Date(2022, 3, 14, 0, 0, 0, 0, time.UTC),
					CreatedAt:     time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:     time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
		},
		"success: empty": {
			givenLimit: 3,
			givenID:    1,
			expResult:  []model.Order{},
		},
		"success: lastID is 101 ": {
			givenLimit: 3,
			givenID:    101,
			expResult: []model.Order{
				{
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
				{
					Id:            99,
					Name:          "an",
					Description:   "an",
					TotalPrice:    99,
					Quantity:      99,
					Discount:      9,
					Shipping:      "hcm",
					Status:        model.StatusShipped,
					UserId:        99,
					DatePurchased: time.Date(2022, 3, 14, 0, 0, 0, 0, time.UTC),
					CreatedAt:     time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					UpdatedAt:     time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
				},
			},
		},
		"success: givenLimit 2": {
			givenLimit: 2,
			givenID:    0,
			expResult: []model.Order{
				{
					Id:            101,
					Name:          "abc1",
					Description:   "abc1",
					TotalPrice:    101,
					Quantity:      101,
					Discount:      11,
					Shipping:      "abc",
					Status:        model.StatusApproved,
					UserId:        101,
					DatePurchased: time.Date(2022, 3, 16, 0, 0, 0, 0, time.UTC),
					CreatedAt:     time.Date(2022, 3, 16, 16, 0, 0, 0, time.UTC),
					UpdatedAt:     time.Date(2022, 3, 16, 16, 0, 0, 0, time.UTC),
				},
				{
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
			},
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
			rs, err := instance.GetOrders(context.Background(), tc.givenLimit, tc.givenID)

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
