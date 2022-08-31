package user

import (
	"context"
	"fmt"
	"order-mg/internal/model"
	"order-mg/internal/util"
)

func (i impl) UpdateUser(ctx context.Context, user model.Users, userID int64) (model.Users, error) {

	userF, err := i.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		fmt.Printf("error when get user by id, userId: %d", userID)
		return model.Users{}, err
	}

	user.Password = util.HashPassword(user.Password)
	comparePassword := util.ComparePassword(userF.Password, user.Password)

	if !comparePassword {
		userF.Password = user.Password
	}
	if userF.Name != user.Name {
		userF.Name = user.Name
	}
	if userF.Address != user.Address {
		userF.Address = user.Address
	}
	if userF.PhoneNumber != user.PhoneNumber {
		userF.PhoneNumber = user.PhoneNumber
	}
	if userF.Age != user.Age {
		userF.Age = user.Age
	}
	i.userRepo.UpdateUser(ctx, userF)

	userF.Password = ""
	return userF, nil
}
