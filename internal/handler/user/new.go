package user

import (
	userSvc "order-mg/internal/service/user"
)

// UserHandler
type UserHandler struct {
	UserSvc userSvc.UserService
}

// New
func New(userService userSvc.UserService) UserHandler {
	return UserHandler{
		UserSvc: userService,
	}
}
