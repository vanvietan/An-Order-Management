package user

import (
	"context"
	"errors"
	"order-mg/internal/model"

	"gorm.io/gorm"
)

// DeleteUser delete a user by id
func (i impl) DeleteUser(ctx context.Context, userID int64) error {
	var tx *gorm.DB
	if tx = i.gormDB.Delete(&model.Users{}, userID); tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected != 1 {
		return errors.New("record not found")
	}
	return nil
}
