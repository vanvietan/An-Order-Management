package user

import "order-mg/internal/model"

type getUsersResponse struct {
	Users  []model.Users `json:"users"`
	Cursor int64         `json:"cursor"`
}
