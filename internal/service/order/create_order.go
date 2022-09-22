package order

import (
	"context"
	log "github.com/sirupsen/logrus"
	"order-mg/internal/model"
	"order-mg/internal/util"
)

var getNextIDFunc = util.GetNextId

// CreateOrder create an order
func (i impl) CreateOrder(ctx context.Context, input model.Order) (model.Order, error) {
	ID, err := getNextIDFunc()
	if err != nil {
		log.Printf("error when generate ID %v ", err)
		return model.Order{}, err
	}
	input.Id = ID

	order, errO := i.orderRepo.CreateOrder(ctx, input)
	if errO != nil {
		log.Printf("error when create an order: %+v", input)
		return model.Order{}, errO
	}
	return order, nil
}
