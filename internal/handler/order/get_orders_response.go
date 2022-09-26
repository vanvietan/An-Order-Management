package order

import "order-mg/internal/model"

type getOrdersResponse struct {
	Orders []model.Order `json:"orders"`
	Cursor int64         `json:"cursor"`
}

func toGetOrdersResponse(orders []model.Order) getOrdersResponse {
	if len(orders) == 0 {
		return getOrdersResponse{}
	}
	return getOrdersResponse{
		Orders: orders,
		Cursor: orders[len(orders)-1].Id,
	}
}

type getOrdersResponseB struct {
	Orders []AnOrderResponse `json:"orders"`
	Cursor int64             `json:"cursor"`
}

func toGetOrdersResponseB(orders []model.Order) getOrdersResponseB {
	if len(orders) == 0 {
		return getOrdersResponseB{}
	}
	var res []AnOrderResponse

	return getOrdersResponseB{
		Orders: res,
		Cursor: res[len(res)-1].Id,
	}
}
