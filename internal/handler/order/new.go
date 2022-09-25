package order

import "order-mg/internal/service/order"

// OrderHandler
type OrderHandler struct {
	OrderSvc order.OrderService
}

// New
func New(orderService order.OrderService) OrderHandler {
	return OrderHandler{
		OrderSvc: orderService,
	}
}
