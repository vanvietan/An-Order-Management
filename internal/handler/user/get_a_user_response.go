package user

import "order-mg/internal/model"

type getAUserResponse struct {
	User   model.Users `json:"user"`
	Cursor int64       `json:"cursor"`
}
