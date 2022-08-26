package user

import (
	"context"
	"fmt"
	"order-mg/internal/model"
)

// CreateUser create a user
func (i impl) CreateUser(ctx context.Context, user model.Users) (model.Users, error) {
	user, err := i.userRepo.CreateUser(ctx, user)
	if err != nil {
		fmt.Errorf("error when get a user, %v", user.Name)
		return model.Users{}, err
	}
	return user, nil

}
