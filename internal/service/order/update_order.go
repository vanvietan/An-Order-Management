package order

import (
	"context"
	log "github.com/sirupsen/logrus"
	"order-mg/internal/model"
)

// UpdateOrder update an order
func (i impl) UpdateOrder(ctx context.Context, input model.Order, orderID int64) (model.Order, error) {
	orderF, err := i.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		log.Printf("error when find an order by ID: %d", orderID)
		return model.Order{}, err
	}

	orderF.Name = input.Name
	orderF.Description = input.Description
	orderF.Discount = input.Discount
	orderF.TotalPrice = input.TotalPrice
	orderF.Quantity = input.Quantity
	orderF.Shipping = input.Shipping
	orderF.Status = input.Status
	orderF.UserId = input.UserId

	orderU, errU := i.orderRepo.UpdateOrder(ctx, orderF)
	if errU != nil {
		log.Printf("error when save order %+v", input)
	}
	return orderU, nil
}
