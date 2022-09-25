package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocks "order-mg/internal/mocks/repository/order"
	"order-mg/internal/model"
	"order-mg/internal/util"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	type createOrder struct {
		mockIn  model.Order
		mockOut model.Order
		mockErr error
	}
	type arg struct {
		givenIn     model.Order
		createOrder createOrder
		expRs       model.Order
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenIn: model.Order{
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
			createOrder: createOrder{
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
				mockOut: model.Order{
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
		"fail: generate id fail": {
			givenIn: model.Order{
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
			createOrder: createOrder{
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
				mockOut: model.Order{},
				mockErr: errors.New("something wrong"),
			},
			expRs:  model.Order{},
			expErr: errors.New("something wrong"),
		},
		"fail: create order fail": {
			givenIn: model.Order{
				Id:            100,
				Name:          "abc",
				Description:   "abc",
				TotalPrice:    100,
				Quantity:      100,
				Discount:      10,
				Shipping:      "abc",
				Status:        model.StatusApproved,
				UserId:        300,
				DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
				CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
			},
			createOrder: createOrder{
				mockIn: model.Order{
					Id:            100,
					Name:          "abc",
					Description:   "abc",
					TotalPrice:    100,
					Quantity:      100,
					Discount:      10,
					Shipping:      "abc",
					Status:        model.StatusApproved,
					UserId:        300,
					DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
					CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
					UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
				},
				mockOut: model.Order{},
				mockErr: errors.New("something wrong"),
			},
			expRs:  model.Order{},
			expErr: errors.New("something wrong"),
		},
	}

	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.OrderRepository)
			instance.On("CreateOrder", mock.Anything, tc.createOrder.mockIn).
				Return(tc.createOrder.mockOut, tc.createOrder.mockErr)

			getNextIDFunc = func() (int64, error) {
				if s == "fail: generate id fail" {
					return 0, errors.New("something wrong")
				}
				return 100, nil
			}
			defer func() {
				getNextIDFunc = util.GetNextId
			}()

			//WHEN
			svc := New(instance)
			rs, err := svc.CreateOrder(context.Background(), tc.givenIn)

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
