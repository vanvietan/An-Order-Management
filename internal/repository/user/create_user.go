package user

import (
	"context"
	"order-mg/internal/model"

	"gorm.io/gorm"
)

// CreateUser create a user
func (i impl) CreateUser(ctx context.Context, user *model.Users) error {
	var tx *gorm.DB
	if user.Username != "" {
		tx = i.gormDB.Select("Name", "Username", "Password", "PhoneNumber", "Address", "Age", "Role").Create(&user)
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
