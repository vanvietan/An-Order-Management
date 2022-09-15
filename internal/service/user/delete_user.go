package user

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

// DeleteUser delete a user
func (i impl) DeleteUser(ctx context.Context, userID int64) error {
	isSuccess, err := i.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		log.Printf("error when deleting a user with id:  %v", userID)
	}
	if !isSuccess {
		return errors.New("can't delete a user")
	}

	return nil
}
