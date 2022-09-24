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

func TestGetOrderByID(t *testing.T) {
	type getOrderByID struct {
		mockID   int64
		mockResp model.Order
		mockErr  error
	}
	type arg struct {
		givenID      int64
		getOrderByID getOrderByID
		expResult    model.Order
		expErr       error
	}

	tcs := map[string]arg{
		"success: get an order": {
			givenID: 100,
			getOrderByID: getOrderByID{
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
			expResult: model.Order{
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
		"fail: can't find ID": {
			givenID: 200,
			getOrderByID: getOrderByID{
				mockID:   200,
				mockResp: model.Order{},
				mockErr:  errors.New("something wrong"),
			},
			expResult: model.Order{},
			expErr:    errors.New("something wrong"),
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.OrderRepository)
			instance.On("GetOrderByID", mock.Anything, tc.getOrderByID.mockID).
				Return(tc.getOrderByID.mockResp, tc.getOrderByID.mockErr)

			//WHEN
			svc := New(instance)
			rs, err := svc.GetOrderByID(context.Background(), tc.givenID)

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
