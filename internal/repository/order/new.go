package order

import (
	"context"
	"gorm.io/gorm"
	"order-mg/internal/model"
)

// OrderRepository contain all order repository functions
type OrderRepository interface {
	//GetOrders get all orders
	GetOrders(ctx context.Context, limit int, lastID int64) ([]model.Order, error)

	//GetOrderByID get a order
	GetOrderByID(ctx context.Context, orderID int64) (model.Order, error)

	//CreateOrder create a order
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)

	//DeleteOrder delete a order
	DeleteOrder(ctx context.Context, orderID int64) error

	//UpdateOrder update a order
	UpdateOrder(ctx context.Context, order model.Order) (model.Order, error)
}
type impl struct {
	gormDB *gorm.DB
}

// New
func New(gormDB *gorm.DB) OrderRepository {
	return impl{
		gormDB: gormDB,
	}
}
