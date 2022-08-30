package user

import (
	"context"
	"errors"
	"order-mg/internal/model"

	"gorm.io/gorm"
)

// UpdateUser: modify user fields
func (i impl) UpdateUser(ctx context.Context, user model.Users, userID int64) (model.Users, error) {
	var tx *gorm.DB

	// if userID == 0 {
	// 	return model.Users{}, tx.Error
	// } else if user.Password == "" {
	// 	return model.Users{}, tx.Error
	// } else {
	// 	tx = i.gormDB.Model(&user).Where("id = ?", userID).Select("name", "password", "phone_number", "address", "age").Updates(map[string]interface{}{"name": user.Name, "password": user.Password, "phone_number": user.PhoneNumber, "address": user.Address, "age": user.Age})
	// }

	if userID == 0 {
		return model.Users{}, errors.New("userID invalid")
	}
	if user.Password == "" {
		return model.Users{}, errors.New("password is empty")
	}

	// tx.

	if tx.Error != nil {
		return model.Users{}, tx.Error
	}

	return user, nil
}
