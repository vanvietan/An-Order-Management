package user

import (
	"context"
	"fmt"
	"order-mg/internal/model"
	"order-mg/internal/util"
)

// CreateUser create a user
func (i impl) CreateUser(ctx context.Context, input model.Users) (model.Users, error) {

	id, err := util.GetNextId()
	if err != nil {
		fmt.Printf("error when generate, %v", err)
		return model.Users{}, err
	}
	input.Id = id

	input.Password = util.HashPassword(input.Password)
	user, errs := i.userRepo.CreateUser(ctx, input)
	if errs != nil {

		fmt.Printf("error when get a user, %v", user.Name)

		return model.Users{}, errs
	}
	user.Password = ""
	return user, nil
}
