package model

import (
	"gorm.io/gorm"
	"time"
)

// Order with entity related to order when user buy it
type Order struct {
	Id            int64          `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	TotalPrice    int32          `json:"total_price"`
	Quantity      int            `json:"quantity"`
	Discount      int8           `json:"discount"`
	Shipping      string         `json:"shipping"`
	Status        Status         `json:"status"` // Enums Status
	UserId        int64          `json:"userID"`
	DatePurchased time.Time      `json:"date_purchased"`
	CreatedAt     time.Time      `json:"created_at" `
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at,omitempty"`
	Histories     []History      `gorm:"foreignKey:OrderId;references:Id" json:"histories,omitempty"`
}
