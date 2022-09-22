package order

import (
	"context"
	log "github.com/sirupsen/logrus"
	"order-mg/internal/model"
)

// GetOrders get all orders
func (i impl) GetOrders(ctx context.Context, limit int, lastID int64) ([]model.Order, error) {
	orders, err := i.orderRepo.GetOrders(ctx, limit, lastID)
	if err != nil {
		log.Printf("error when get orders, limit %d, last %d", limit, lastID)
		return nil, err
	}
	return orders, nil
}
