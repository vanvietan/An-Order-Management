package user

import (
	"context"
	"order-mg/internal/model"
)

// GetUserByID find a user by its id
func (i impl) GetUserByID(ctx context.Context, userID int64) (model.Users, error) {
	user := model.Users{}
	tx := i.gormDB.First(&user, userID)
	if tx.Error != nil {
		return model.Users{}, tx.Error
	}
	return user, nil
}
