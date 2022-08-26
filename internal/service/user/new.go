package user

import (
	"context"
	"order-mg/internal/model"
	"order-mg/internal/repository/user"
)

// UserService: handle all user service busineess
type UserService interface {
	//GetUsers get all users
	GetUsers(ctx context.Context, limit int, lastID int64) ([]model.Users, error)

	//GetUserByID find a user by its id
	GetUserByID(ctx context.Context, userId int64) (model.Users, error)

	// CreateUser create a user
	CreateUser(ctx context.Context, user model.Users) (model.Users, error)
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
