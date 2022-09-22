package order

import (
	"context"
	"gorm.io/gorm"
	"order-mg/internal/model"
)

// GetOrders get all orders
func (i impl) GetOrders(ctx context.Context, limit int, lastID int64) ([]model.Order, error) {
	orders := make([]model.Order, limit)
	var tx *gorm.DB
	if lastID == 0 {
		tx = i.gormDB.Model(model.Order{}).Select("orders.*").Limit(limit).Order("created_at desc").Find(&orders)
	} else {
		tx = i.gormDB.Select("orders.*").Where("id < ?", lastID).Limit(limit).Order("created_at desc").Find(&orders)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}
	return orders, nil
}
