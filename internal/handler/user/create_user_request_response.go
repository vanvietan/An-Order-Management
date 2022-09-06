package user

import (
	"errors"
	"order-mg/internal/model"
	"time"
)

// CreateUserInput input from clients
type CreateUserInput struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Username    string     `json:"username"`
	Password    string     `json:"password,omitempty"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	Age         int8       `json:"age"`
	Role        model.Role `json:"role"` //Enums Role
}

// AUserResponse response to json
type AUserResponse struct {
	Id          int64           `json:"id"`
	Name        string          `json:"name"`
	Username    string          `json:"username"`
	PhoneNumber string          `json:"phone_number"`
	Address     string          `json:"address"`
	Age         int8            `json:"age"`
	Role        model.Role      `json:"role"` //Enums Role
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Orders      []model.Order   `json:"Orders,omitempty"`
	Histories   []model.History `json:"Histories,omitempty"`
}

func (c CreateUserInput) validateAndMap() (model.Users, error) {
	if c.Name == "" {
		return model.Users{}, errors.New("invalid name")
	}
	if c.Username == "" || len(c.Username) > 14 {
		return model.Users{}, errors.New("invalid username")
	}
	if c.Password == "" || len(c.Password) > 14 {
		return model.Users{}, errors.New("password is invalid")
	}
	if c.PhoneNumber == "" || len(c.PhoneNumber) > 11 || len(c.PhoneNumber) < 10 {
		return model.Users{}, errors.New("phone number is invalid")
	}
	if c.Address == "" || len(c.Address) > 120 {
		return model.Users{}, errors.New("address is invalid")
	}
	if c.Age <= 0 || c.Age > 120 {
		return model.Users{}, errors.New("user age is invalid")
	}
	if c.Role != model.RoleUser {
		return model.Users{}, errors.New("user role is invalid")
	}
	return model.Users{
		Name:        c.Name,
		Username:    c.Username,
		Password:    c.Password,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
		Age:         c.Age,
		Role:        c.Role,
	}, nil
}

func toGetAUserResponse(user model.Users) AUserResponse {
	return AUserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Age:         user.Age,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Orders:      user.Orders,
		Histories:   user.Histories,
	}
}
