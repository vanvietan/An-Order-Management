package model

import (
	"time"
)

// History(Audit_Trail) tracking operation by admin depends on
type History struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	OrderId   int64     `json:"orderId"`
	Operation Operation `json:"operation"` //Enums Operation
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
