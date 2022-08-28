package user

import (
	"context"
	"fmt"
)

// DeleteUser delete a user
func (i impl) DeleteUser(ctx context.Context, lastID int64) (bool, error) {

	isSucess, err := i.userRepo.DeleteUser(ctx, lastID)
	if err != nil {
		fmt.Errorf("error when deleting a user with id , %v", lastID)
	}
	if !isSucess {
		return false, err
	}

	return true, nil
}
