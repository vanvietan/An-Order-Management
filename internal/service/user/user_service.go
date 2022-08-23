package user

import (
	"context"
	"order-mg/internal/model"
)

// GetUsers get all users
func (i impl) GetUsers(ctx context.Context, limit int, lastID int64) ([]model.Users, error) {
	return i.userRepo.GetUsers(ctx, limit, lastID)
}
