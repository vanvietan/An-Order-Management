package model

import (
	"time"
)

type Order struct {
	Id            uint         `json:"id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	TotalPrice    int32        `json:"totalPrice"`
	Quantity      int          `json:"quantity"`
	Discount      int8         `json:"discount"`
	Shipping      string       `json:"shipping"`
	Status        int          `json:"status"`
	UserId        int          `json:"userId"`
	DatePurchased time.Weekday `json:"datePurchased"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
	DeletedAt     time.Time    `json:"deletedAt"`
}

type Users struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
	Age         int8      `json:"age"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type History struct {
	Id        uint      `json:"id"`
	UserId    int       `json:"userId"`
	OrderId   int       `json:"orderId"`
	Operation int       `json:"operation"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
