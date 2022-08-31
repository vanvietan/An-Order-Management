package user

import (
	"context"
	"fmt"
	"order-mg/internal/model"
	"order-mg/internal/util"
)

// CreateUser create a user
func (i impl) CreateUser(ctx context.Context, user model.Users) (model.Users, error) {

	id, err := util.GetNextId()
	if err != nil {
		fmt.Printf("error when generate, %v", err)
		return model.Users{}, err
	}
	user.Id = id

	user.Password = util.HashPassword(user.Password)
	_, errs := i.userRepo.CreateUser(ctx, user)
	if errs != nil {

		fmt.Printf("error when get a user, %v", user.Name)

		return model.Users{}, errs
	}
	user.Password = ""
	return user, nil
}
