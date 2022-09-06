package user

import (
	"context"
	"order-mg/internal/model"
	"order-mg/internal/util"

	log "github.com/sirupsen/logrus"
)

// CreateUser create a user
func (i impl) CreateUser(ctx context.Context, input model.Users) (model.Users, error) {

	id, err := util.GetNextId()
	if err != nil {
		log.Printf("error when generate, %v", err)
		return model.Users{}, err
	}
	input.Id = id

	input.Password = util.HashPassword(input.Password)
	user, errs := i.userRepo.CreateUser(ctx, input)
	if errs != nil {

		log.Printf("error when get a user, %v", user.Name)

		return model.Users{}, errs
	}
	user.Password = ""
	return user, nil
}
