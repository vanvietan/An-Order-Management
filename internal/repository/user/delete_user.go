package user

import (
	"context"
	"order-mg/internal/model"

	"gorm.io/gorm"
)

// DeleteUser delete a user by id
func (i impl) DeleteUser(ctx context.Context, userID int64) (bool, error) {
	var tx *gorm.DB

	user := model.Users{}

	tx = i.gormDB.Delete(&user, userID)

	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
