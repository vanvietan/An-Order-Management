package handler

import (
	"order-mg/internal/handler/order"
	"order-mg/internal/handler/user"
	orderService "order-mg/internal/service/order"
	userService "order-mg/internal/service/user"
)

// Handler
type Handler struct {
	UserHandler  user.UserHandler
	OrderHandler order.OrderHandler
}

// New
func New(userSvc userService.UserService, orderSvc orderService.OrderService) Handler {
	return Handler{
		UserHandler: user.UserHandler{
			UserSvc: userSvc,
		},
		OrderHandler: order.OrderHandler{
			OrderSvc: orderSvc,
		},
	}
}
