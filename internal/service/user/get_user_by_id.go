package user

import (
	"context"
	"errors"
	"order-mg/internal/model"

	log "github.com/sirupsen/logrus"
)

// GetUserByID find a user by id
func (i impl) GetUserByID(ctx context.Context, userID int64) (model.Users, error) {
	if userID <= 0 {
		return model.Users{}, errors.New("invalid userID")
	}
	user, err := i.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Printf("error when get user by id, userId: %d", userID)
		return model.Users{}, err
	}
	user.Password = ""
	return user, nil

}
