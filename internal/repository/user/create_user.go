package user

import (
	"context"
	"order-mg/internal/model"

	"gorm.io/gorm"
)

// CreateUser create a user
func (i impl) CreateUser(ctx context.Context, user model.Users) (model.Users, error) {
	var tx *gorm.DB
	if user.Username == "" {
		return model.Users{}, tx.Error
	} else if user.Password == "" {
		return model.Users{}, tx.Error
	} else {
		tx = i.gormDB.Create(&user)
	}
	if tx.Error != nil {
		return model.Users{}, tx.Error
	}
	return user, nil
}
