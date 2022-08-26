package router

import (
	"order-mg/internal/handler"
	userSvc "order-mg/internal/service/user"

	"github.com/go-chi/chi"
)

// MasterRoute
type MasterRoute struct {
	Router      *chi.Mux
	Handler     handler.Handler
	UserService userSvc.UserService
}

func New(r *chi.Mux, userService userSvc.UserService) {
	handler := handler.New(userService)
	mr := MasterRoute{
		Router:  r,
		Handler: handler,
	}
	mr.initRoutes()
}

func (mr MasterRoute) initRoutes() {
	mr.initUserRouters()
}

func (mr MasterRoute) initUserRouters() {
	mr.Router.Group(func(r chi.Router) {
		r.Get("/users", mr.Handler.UserHandler.GetUsers)
		r.Get("/users/{cursor}", mr.Handler.UserHandler.GetUserByID)
	})
}
