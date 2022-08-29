package user

import (
	"context"
	"order-mg/internal/model"

	"gorm.io/gorm"
)

// UserRepository contain all user repository functions
type UserRepository interface {
	// GetUsers get all users
	GetUsers(ctx context.Context, limit int, lastID int64) ([]model.Users, error)

	//GetUserById: find a user by its id
	GetUserByID(ctx context.Context, userID int64) (model.Users, error)

	// CreateUser create a user
	CreateUser(ctx context.Context, user model.Users) (model.Users, error)

	//DeleteUser delete a user
	DeleteUser(ctx context.Context, userID int64) (bool, error)
}

type impl struct {
	gormDB *gorm.DB
}

func New(gormDB *gorm.DB) impl {
	return impl{
		gormDB: gormDB,
	}
}
