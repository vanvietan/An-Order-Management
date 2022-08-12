package model

import (
	_ "github.com/lib/pq"
	"time"
)

type Order struct {
	Id            int          `json:"id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	totalPrice    float32      `json:"total_Price"`
	quantity      int          `json:"quantity"`
	discount      int          `json:"discount"`
	shipping      string       `json:"shipping"`
	paymentMethod string       `json:"payment_Method"`
	orderStatus   int          `json:"order_Status"`
	userId        int          `json:"userId"`
	datePurchased time.Weekday `json:"datePurchased"`
	createdAt     time.Time    `json:"createdAt"`
	updatedAt     time.Time    `json:"updatedAt"`
}
