package user

import "order-mg/internal/model"

type getUsersResponse struct {
	Users  []AUserResponse `json:"users"`
	Cursor int64           `json:"cursor"`
}

func toGetUsersResponse(resp []AUserResponse) getUsersResponse {
	if len(resp) == 0 {
		return getUsersResponse{}
	}

	return getUsersResponse{
		Users:  resp,
		Cursor: resp[len(resp)-1].Id,
	}
}

func modelToResponseArray(users []model.Users) []AUserResponse {
	if len(users) == 0 {
		return nil
	}
	resp := make([]AUserResponse, len(users))
	for i, s := range users {
		resp[i].Id = s.Id
		resp[i].Name = s.Name
		resp[i].Username = s.Username
		resp[i].PhoneNumber = s.PhoneNumber
		resp[i].Address = s.Address
		resp[i].Age = s.Age
		resp[i].Role = s.Role
		resp[i].CreatedAt = s.CreatedAt
		resp[i].UpdatedAt = s.UpdatedAt
		resp[i].Orders = s.Orders
		resp[i].Histories = s.Histories
	}
	return resp
}
