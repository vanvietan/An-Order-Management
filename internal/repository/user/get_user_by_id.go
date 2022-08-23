package user

import (
	"context"
	"order-mg/internal/model"

	"gorm.io/gorm"
)

// GetUserById: find a user by its id
func (i impl) GetUserById(ctx context.Context, userId int64) (model.Users, error) {
	user := model.Users{}
	var tx *gorm.DB
	if userId == 0 {
		return model.Users{}, tx.Error
	} else if userId < 0 {
		return model.Users{}, tx.Error
	} else {
		tx = i.gormDB.First(&user, userId)
	}
	if tx.Error != nil {
		return model.Users{}, tx.Error
	}
	return user, nil
}
