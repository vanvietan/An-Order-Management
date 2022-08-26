package user

import (
	"context"
	"fmt"
	"order-mg/internal/model"
)

// GetUserByID find a user by id
func (i impl) GetUserByID(ctx context.Context, userID int64) (model.Users, error) {
	user, err := i.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		fmt.Errorf("error when get user by id, userId: %d", userID)
		return model.Users{}, err
	}

	return user, nil

}
