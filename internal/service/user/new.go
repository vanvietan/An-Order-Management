package user

import (
	"context"
	"order-mg/internal/model"
	"order-mg/internal/repository/user"
)

// UserService: handle all user service busineess
type UserService interface {
	//GetUsers get all users
	GetUsers(ctx context.Context, limit int, lastID int64) (error, []model.Users)
}
type impl struct {
	userRepo user.UserRepository
}

// New
func New(userRepo user.UserRepository) impl {
	return impl{
		userRepo: userRepo,
	}
}
