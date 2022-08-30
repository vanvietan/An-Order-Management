package user

import (
	"context"
	"order-mg/internal/model"
)

// CreateUser create a user
func (i impl) CreateUser(ctx context.Context, user model.Users) (model.Users, error) {

	tx := i.gormDB.Create(&user)
	if tx.Error != nil {
		return model.Users{}, tx.Error
	}
	return user, nil
}
