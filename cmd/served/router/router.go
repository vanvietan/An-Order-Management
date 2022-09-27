package router

import (
	"github.com/go-chi/chi"
	"order-mg/internal/handler"
	orderSvc "order-mg/internal/service/order"
	userSvc "order-mg/internal/service/user"
)

// MasterRoute
type MasterRoute struct {
	Router       *chi.Mux
	Handler      handler.Handler
	UserService  userSvc.UserService
	OrderService orderSvc.OrderService
}

func New(r *chi.Mux, userService userSvc.UserService, orderService orderSvc.OrderService) {
	handler := handler.New(userService, orderService)
	mr := MasterRoute{
		Router:  r,
		Handler: handler,
	}
	mr.initRoutes()
}

func (mr MasterRoute) initRoutes() {
	mr.initUserRouters()
	mr.initOrderRouters()
}

func (mr MasterRoute) initUserRouters() {
	mr.Router.Group(func(r chi.Router) {
		r.Get("/users", mr.Handler.UserHandler.GetUsers)
		r.Get("/users/{id}", mr.Handler.UserHandler.GetUserByID)
		r.Post("/users", mr.Handler.UserHandler.CreateUser)
		r.Delete("/users/{id}", mr.Handler.UserHandler.DeleteUser)
		r.Put("/users/{id}", mr.Handler.UserHandler.UpdateUser)
	})
}
func (mr MasterRoute) initOrderRouters() {
	mr.Router.Group(func(r chi.Router) {
		r.Get("/orders", mr.Handler.OrderHandler.GetOrders)
		r.Get("/orders/{id}", mr.Handler.OrderHandler.GetOrderByID)
		r.Post("/orders", mr.Handler.OrderHandler.CreateOrder)
		r.Put("/orders/{id}", mr.Handler.OrderHandler.UpdateOrder)
		r.Delete("/orders/{id}", mr.Handler.OrderHandler.DeleteOrder)
	})
}
