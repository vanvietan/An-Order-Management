package user

import (
	"order-mg/internal/model"
	"time"
)

type Input struct {
	Id          int64
	Name        string
	Username    string
	Password    string
	PhoneNumber string
	Address     string
	Age         int8
	Role        model.Role //Enums Role
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
