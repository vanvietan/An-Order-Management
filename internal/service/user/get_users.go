package user

import (
	"context"
	"order-mg/internal/model"

	log "github.com/sirupsen/logrus"
)

// GetUsers get all users
func (i impl) GetUsers(ctx context.Context, limit int, lastID int64) ([]model.Users, error) {
	users, err := i.userRepo.GetUsers(ctx, limit, lastID)
	if err != nil {
		log.Printf("error when get users, limit: %d, lastID: %d", limit, lastID)
		return nil, err
	}
	for i := range users {
		users[i].Password = ""
	}
	return users, nil
}
