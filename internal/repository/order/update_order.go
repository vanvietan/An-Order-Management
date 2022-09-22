package order

import (
	"context"
	"order-mg/internal/model"
)

// UpdateOrder update an order
func (i impl) UpdateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	tx := i.gormDB.Save(&order)
	if tx.Error != nil {
		return model.Order{}, tx.Error
	}
	return order, nil
}
