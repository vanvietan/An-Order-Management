package user

import (
	"context"
	"fmt"
	"order-mg/internal/model"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 10
)

// CreateUser create a user
func (i impl) CreateUser(ctx context.Context, user model.Users) (model.Users, error) {
	user.Password = hashPassword(user.Password)
	_, err := i.userRepo.CreateUser(ctx, user)
	if err != nil {
		fmt.Errorf("error when get a user, %v", user.Name)
		return model.Users{}, err
	}
	return user, nil

}

func hashPassword(s string) string {
	bs, err := bcrypt.GenerateFromPassword([]byte(s), MinCost)
	if err != nil {
		fmt.Errorf("error when hash pw, %v", bs)
		return ""
	}
	bpass := string(bs)
	return bpass
}
