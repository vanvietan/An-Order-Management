package user

import (
	"context"

	log "github.com/sirupsen/logrus"
)

// DeleteUser delete a user
func (i impl) DeleteUser(ctx context.Context, userID int64) error {
	err := i.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		log.Printf("error when deleting user :  %v", err)
	}

	return err
}
