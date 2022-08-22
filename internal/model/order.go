package model

import (
	"time"
)

// Order: with entity related to order when user buy it
type Order struct {
	Id            int64        `json:"id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	TotalPrice    int32        `json:"totalPrice"`
	Quantity      int          `json:"quantity"`
	Discount      int8         `json:"discount"`
	Shipping      string       `json:"shipping"`
	Status        Status       `json:"status"` // Enums Status
	UserId        int64        `json:"userId"`
	DatePurchased time.Weekday `json:"datePurchased"`
	CreatedAt     time.Time    `json:"created_at" `
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     time.Time    `json:"deleted_at"`
	Histories     []History    `gorm:"foreignKey:OrderId;references:Id"`
}
