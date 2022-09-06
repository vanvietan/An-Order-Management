package user

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

// DeleteUser delete a user
func (i impl) DeleteUser(ctx context.Context, userID int64) (bool, error) {
	if userID <= 0 {
		return false, errors.New("invalid userID")
	}
	_, err := i.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return false, errors.New("userID is not existed")
	}
	isSucess, err := i.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		log.Printf("error when deleting a user with id , %v", userID)
	}
	if !isSucess {
		return false, errors.New("can't delete a user")
	}

	return true, nil
}
