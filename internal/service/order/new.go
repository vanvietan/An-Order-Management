package order

import (
	"context"
	"order-mg/internal/model"
	"order-mg/internal/repository/order"
)

// OrderService handle all order service business
type OrderService interface {
	//GetOrders get all orders
	GetOrders(ctx context.Context, limit int, lastID int64) ([]model.Order, error)

	//GetOrderByID get an order
	GetOrderByID(ctx context.Context, orderID int64) (model.Order, error)

	//CreateOrder create an order
	CreateOrder(ctx context.Context, input model.Order) (model.Order, error)

	//UpdateOrder update an orders
	UpdateOrder(ctx context.Context, input model.Order, orderID int64) (model.Order, error)

	//DeleteOrder delete an order
	DeleteOrder(ctx context.Context, orderID int64) error
}

type impl struct {
	orderRepo order.OrderRepository
}

// New
func New(orderRepo order.OrderRepository) OrderService {
	return impl{
		orderRepo: orderRepo,
	}
}
