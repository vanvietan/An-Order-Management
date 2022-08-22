package user

import (
	"context"
	"order-mg/internal/model"

	"gorm.io/gorm"
)

// GetUsers get all users
func (i impl) GetUsers(ctx context.Context, limit int, lastID int64) ([]model.Users, error) {
	users := make([]model.Users, limit)
	var tx *gorm.DB
	if lastID == 0 {
		tx = i.gormDB.Model(model.Users{}).Select("users.*").Limit(limit).Order("created_at desc").Find(&users)
	} else {
		tx = i.gormDB.Select("users.*").Where("id < ?", lastID).Limit(limit).Order("created_at desc").Find(&users)
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
