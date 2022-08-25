package user

import (
	"context"
	"order-mg/internal/model"

	"gorm.io/gorm"
)

// GetUserById: find a user by its id
func (i impl) GetUserByID(ctx context.Context, userID int64) (model.Users, error) {
	user := model.Users{}
	var tx *gorm.DB
	if userID == 0 {
		return model.Users{}, tx.Error
	} else if userID < 0 {
		return model.Users{}, tx.Error
	} else {
		tx = i.gormDB.First(&user, userID)
	}
	if tx.Error != nil {
		return model.Users{}, tx.Error
	}
	return user, nil
}
