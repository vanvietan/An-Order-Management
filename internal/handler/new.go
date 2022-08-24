package handler

import (
	"order-mg/internal/handler/user"
	userService "order-mg/internal/service/user"
)

// Handler
type Handler struct {
	UserHandler user.UserHandler
}

// New
func New(userSvc userService.UserService) Handler {
	return Handler{
		UserHandler: user.UserHandler{
			UserSvc: userSvc,
		},
	}
}
