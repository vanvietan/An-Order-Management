package model

import (
	"time"
)

// Users with 2 Role User/Admin to create or modified order
type Users struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Password    string    `json:"password,omitempty"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Age         int       `json:"age"`
	Role        Role      `json:"role"` //Enums Role
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Orders      []Order   `gorm:"foreignKey:UserId;references:Id" json:"Orders,omitempty"`
	Histories   []History `gorm:"foreignKey:UserId;references:Id" json:"Histories,omitempty"`
}
