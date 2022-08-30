package user

import (
	"context"
	"order-mg/internal/model"
)

// UpdateUser: modify user fields
func (i impl) UpdateUser(ctx context.Context, user model.Users) (model.Users, error) {

	tx := i.gormDB.Save(&user)

	if tx.Error != nil {
		return model.Users{}, tx.Error
	}

	return user, nil
}
