package order

import (
	"order-mg/internal/model"
)

func modelToResponseArray(orders []model.Order) []AnOrderResponse {
	if len(orders) == 0 {
		return nil
	}
	resp := make([]AnOrderResponse, len(orders))
	for i, s := range orders {
		resp[i].Id = s.Id
		resp[i].Name = s.Name
		resp[i].Description = s.Description
		resp[i].TotalPrice = s.TotalPrice
		resp[i].Quantity = s.Quantity
		resp[i].Discount = s.Discount
		resp[i].Shipping = s.Shipping
		resp[i].Status = s.Status
		resp[i].UserId = s.UserId
		resp[i].DatePurchased = s.DatePurchased
		resp[i].CreatedAt = s.CreatedAt
		resp[i].UpdatedAt = s.UpdatedAt
		resp[i].Histories = s.Histories
	}
	return resp
}

type getOrdersResponse struct {
	Orders []AnOrderResponse `json:"orders"`
	Cursor int64             `json:"cursor"`
}

func toGetOrdersResponse(resp []AnOrderResponse) getOrdersResponse {
	if len(resp) == 0 {
		return getOrdersResponse{}
	}
	return getOrdersResponse{
		Orders: resp,
		Cursor: resp[len(resp)-1].Id,
	}
}
