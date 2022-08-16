package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id            uint         `json:"id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	TotalPrice    int32        `json:"totalPrice"`
	Quantity      int          `json:"quantity"`
	Discount      int8         `json:"discount"`
	Shipping      string       `json:"shipping"`
	Status        int          `json:"status"` // Enums Status
	UserId        uint         `json:"userId"`
	DatePurchased time.Weekday `json:"datePurchased"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
	DeletedAt     time.Time    `json:"deletedAt"`
	Histories     []History    `gorm:"foreignKey:OrderId"`
}

type Users struct {
	gorm.Model
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phoneNumber"`
	Address     string    `json:"address"`
	Age         int8      `json:"age"`
	Role        string    `json:"role"` //Enums Role
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Orders      []Order   `gorm:"foreignKey:UserId"`
	Histories   []History `gorm:"foreignKey:UserId"`
}

type History struct {
	gorm.Model
	Id        uint      `json:"id"`
	UserId    uint      `json:"userId"`
	OrderId   uint      `json:"orderId"`
	Operation int       `json:"operation"` //Enums Operation
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Role string
type Status string
type Operation string

const (
	USER  Role = "user"
	ADMIN      = "admin"
)

const (
	APPROVED         Status = "approved"
	APPROVAL_PENDING        = "approval pending"
	SHIPPING                = "shipping"
	SHIPPED                 = "shipped"
)

const (
	MODIFIED Operation = "modified"
	DELETED            = "deleted"
)
