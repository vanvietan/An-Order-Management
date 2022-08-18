package model

import (
	"time"

	"gorm.io/gorm"
)

// Order: with entity related to order when user buy it
type Order struct {
	gorm.Model
	Id            uint           `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	TotalPrice    int32          `json:"totalPrice"`
	Quantity      int            `json:"quantity"`
	Discount      int8           `json:"discount"`
	Shipping      string         `json:"shipping"`
	Status        int            `json:"status"` // Enums Status
	UserId        uint           `json:"userId"`
	DatePurchased time.Weekday   `json:"datePurchased"`
	CreatedAt     time.Time      `json:"created_at" `
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt"`
	Histories     []History      `gorm:"foreignKey:OrderId;references:Id"`
}
