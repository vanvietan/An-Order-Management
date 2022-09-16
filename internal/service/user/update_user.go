package user

import (
	"context"
	"order-mg/internal/model"
	"order-mg/internal/util"

	log "github.com/sirupsen/logrus"
)

var hashPasswordFunc = util.HashPassword

// UpdateUser : modify user
func (i impl) UpdateUser(ctx context.Context, input model.Users, userID int64) (model.Users, error) {

	userF, err := i.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Printf("error when get user by id, userId: %d", userID)
		return model.Users{}, err
	}

	input.Password = hashPasswordFunc(input.Password)
	userF.Password = input.Password
	userF.Name = input.Name
	userF.Address = input.Address
	userF.PhoneNumber = input.PhoneNumber
	userF.Age = input.Age
	userF.Role = input.Role

	userU, errU := i.userRepo.UpdateUser(ctx, userF)
	if errU != nil {
		log.Printf("error when save user, userId: %d", userF.Id)
	}

	userU.Password = ""
	return userU, nil
}
