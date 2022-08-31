package user

import (
	"context"
	"fmt"
	"order-mg/internal/model"
)

// GetUsers get all users
func (i impl) GetUsers(ctx context.Context, limit int, lastID int64) ([]model.Users, error) {
	users, err := i.userRepo.GetUsers(ctx, limit, lastID)
	if err != nil {
		fmt.Printf("error when get users, limit: %d, lastID: %d", limit, lastID)
		return nil, err
	}
	return users, nil
}
