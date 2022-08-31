package user

import (
	"context"
	"errors"
	"fmt"
)

// DeleteUser delete a user
func (i impl) DeleteUser(ctx context.Context, userID int64) (bool, error) {
	if userID <= 0 {
		return false, errors.New("invalid userId")
	}
	isSucess, err := i.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		fmt.Printf("error when deleting a user with id , %v", userID)
	}
	if !isSucess {
		return false, errors.New("can't delete a user")
	}

	return true, nil
}
