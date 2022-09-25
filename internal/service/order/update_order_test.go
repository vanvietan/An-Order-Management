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

func TestUpdateOrder(t *testing.T) {
	type updateOrder struct {
		mockIn   model.Order
		mockID   int64
		mockResp model.Order
		mockErr  error
	}
	type arg struct {
		givenID     int64
		givenIn     model.Order
		updateOrder updateOrder
		expRs       model.Order
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: 100,
			givenIn: model.Order{
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
			updateOrder: updateOrder{
				mockIn: model.Order{
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
				mockID: 100,
				mockResp: model.Order{
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
			expRs: model.Order{
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
		"fail: can't find orderID": {
			givenIn: model.Order{
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
			givenID: 100,
			updateOrder: updateOrder{
				mockIn: model.Order{
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
				mockID:   100,
				mockResp: model.Order{},
				mockErr:  errors.New("something wrong"),
			},
			expRs:  model.Order{},
			expErr: errors.New("something wrong"),
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.OrderRepository)
			instance.On("GetOrderByID", mock.Anything, tc.updateOrder.mockID).
				Return(tc.updateOrder.mockResp, tc.updateOrder.mockErr)
			instance.On("UpdateOrder", mock.Anything, tc.updateOrder.mockIn).
				Return(tc.updateOrder.mockResp, tc.updateOrder.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.UpdateOrder(context.Background(), tc.givenIn, tc.givenID)

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
