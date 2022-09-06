package user

import (
	"context"
	"order-mg/internal/model"
	"order-mg/internal/util"

	log "github.com/sirupsen/logrus"
)

func (i impl) UpdateUser(ctx context.Context, user model.Users, userID int64) (model.Users, error) {

	userF, err := i.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Printf("error when get user by id, userId: %d", userID)
		return model.Users{}, err
	}

	user.Password = util.HashPassword(user.Password)
	userF.Password = user.Password
	userF.Name = user.Name
	userF.Address = user.Address
	userF.PhoneNumber = user.PhoneNumber
	userF.Age = user.Age
	userF.Role = user.Role

	i.userRepo.UpdateUser(ctx, userF)

	userF.Password = ""
	return userF, nil
}
