package order

import (
	"net/http"
	"order-mg/internal/handler/common"
)

// GetOrderByID get an order by ID
func (h OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	orderID, err := validateIDAndMap(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	order, errH := h.OrderSvc.GetOrderByID(r.Context(), orderID)
	if errH != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJson(w, http.StatusOK, toGetAnOrderResponse(order))
}
