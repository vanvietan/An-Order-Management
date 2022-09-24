package order

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mocks "order-mg/internal/mocks/repository/order"
	"testing"
)

func TestDeleteOrder(t *testing.T) {
	type deleteOrder struct {
		mockID  int64
		mockErr error
	}
	type arg struct {
		givenID     int64
		deleteOrder deleteOrder
		expErr      error
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: 99,
			deleteOrder: deleteOrder{
				mockID: 99,
			},
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.OrderRepository)
			instance.On("DeleteOrder", mock.Anything, tc.deleteOrder.mockID).
				Return(tc.deleteOrder.mockErr)

			//WHEN
			svc := New(instance)
			err := svc.DeleteOrder(context.Background(), tc.givenID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
