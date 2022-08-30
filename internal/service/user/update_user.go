package user

import (
	"context"
	"errors"
	"fmt"
	"order-mg/internal/model"
	"order-mg/internal/util"
)

func (i impl) UpdateUser(ctx context.Context, user model.Users, userID int64) (model.Users, error) {

	if userID <= 0 {
		return model.Users{}, errors.New("id is invalid")
	}
	if user.Password == "" {
		return model.Users{}, errors.New("password is empty")
	}
	if user.PhoneNumber == "" {
		return model.Users{}, errors.New("phone number is empty")
	}
	if user.Address == "" {
		return model.Users{}, errors.New("address is empty")
	}
	if user.Age <= 0 || user.Age > 100 {
		return model.Users{}, errors.New("user age is invalid")
	}

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

	return userF, nil
}
