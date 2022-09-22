package order

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"order-mg/internal/model"
)

// DeleteOrder delete an order
func (i impl) DeleteOrder(ctx context.Context, orderID int64) error {
	var tx *gorm.DB
	if tx = i.gormDB.Delete(&model.Order{}, orderID); tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected != 1 {
		return errors.New("record not found")
	}
	return nil
}
