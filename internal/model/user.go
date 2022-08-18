package model

import (
	"time"

	"gorm.io/gorm"
)

// Users with 2 Role User/Admin to create or modified order
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
	Orders      []Order   `gorm:"foreignKey:UserId;references:Id"`
	Histories   []History `gorm:"foreignKey:UserId;references:Id"`
}
