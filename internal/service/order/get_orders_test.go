package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocks "order-mg/internal/mocks/repository/order"
	"order-mg/internal/model"
	"testing"
	"time"
)

func TestGetOrders(t *testing.T) {
	type getOrders struct {
		mockLimit int
		mockID    int64
		mockResp  []model.Order
		mockErr   error
	}
	type arg struct {
		givenLimit  int
		givenLastID int64
		getOrders   getOrders
		expResult   []model.Order
		expErr      error
	}

	tcs := map[string]arg{
		"success: get all orders": {
			givenLimit:  3,
			givenLastID: 0,
			getOrders: getOrders{
				mockLimit: 3,
				mockID:    0,
				mockResp: []model.Order{
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
		"fail: empty result ": {
			givenLimit:  3,
			givenLastID: 1,
			getOrders: getOrders{
				mockLimit: 3,
				mockID:    1,
				mockResp:  []model.Order{},
				mockErr:   errors.New("something wrong"),
			},
			expResult: []model.Order{},
			expErr:    errors.New("something wrong"),
		},
		"success: last id is 101 ": {
			givenLimit:  3,
			givenLastID: 101,
			getOrders: getOrders{
				mockLimit: 3,
				mockID:    101,
				mockResp: []model.Order{
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
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.OrderRepository)
			instance.On("GetOrders", mock.Anything, tc.getOrders.mockLimit, tc.getOrders.mockID).
				Return(tc.getOrders.mockResp, tc.getOrders.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.GetOrders(context.Background(), tc.givenLimit, tc.givenLastID)

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
