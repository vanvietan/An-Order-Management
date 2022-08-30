package user

import (
	"context"
	"errors"
	"log"
	"order-mg/internal/model"
)

// GetUserByID find a user by id
func (i impl) GetUserByID(ctx context.Context, userID int64) (model.Users, error) {
	if userID <= 0 {
		return model.Users{}, errors.New("invalid userID")
	}
	user, err := i.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		log.Fatalf("error when get user by id, userId: %d", userID)
		return model.Users{}, err
	}

	return user, nil

}
