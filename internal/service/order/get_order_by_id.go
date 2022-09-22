package order

import (
	"context"
	log "github.com/sirupsen/logrus"
	"order-mg/internal/model"
)

// GetOrderByID get an order by ID
func (i impl) GetOrderByID(ctx context.Context, orderID int64) (model.Order, error) {
	order, err := i.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		log.Printf("error when get order by id , orderID: %d", orderID)
	}
	return order, nil
}
