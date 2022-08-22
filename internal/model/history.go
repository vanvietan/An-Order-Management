package model

import (
	"time"

	"gorm.io/gorm"
)

// History(Audit_Trail) tracking operation by admin depends on
type History struct {
	gorm.Model
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	OrderId   int64     `json:"orderId"`
	Operation int       `json:"operation"` //Enums Operation
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
