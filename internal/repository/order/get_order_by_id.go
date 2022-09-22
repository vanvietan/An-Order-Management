package order

import (
	"context"
	"order-mg/internal/model"
)

// GetOrderByID get an order by id
func (i impl) GetOrderByID(ctx context.Context, orderID int64) (model.Order, error) {
	order := model.Order{}
	tx := i.gormDB.First(&order, orderID)
	if tx.Error != nil {
		return model.Order{}, tx.Error
	}
	return order, nil
}
