package order

import (
	"encoding/json"
	"net/http"
	"order-mg/internal/handler/common"
	"order-mg/internal/model"
)

// CreateOrder create an order
func (h OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	reqBody, err := checkValidate(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	orderCreated, errH := h.OrderSvc.CreateOrder(r.Context(), reqBody)
	if errH != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	common.ResponseJson(w, http.StatusOK, toGetAnOrderResponse(orderCreated))
}

func checkValidate(r *http.Request) (model.Order, error) {
	var input CreateOrderInput
	_ = json.NewDecoder(r.Body).Decode(&input)
	svcInput, err := input.validateAndMap()
	if err != nil {
		return model.Order{}, err
	}
	return svcInput, nil
}
