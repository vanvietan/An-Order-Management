package order

import (
	"net/http"
	"order-mg/internal/handler/common"
)

// DeleteOrder delete an order
func (h OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := validateIDAndMap(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errD := h.OrderSvc.DeleteOrder(r.Context(), orderID)
	if errD != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	common.ResponseJson(w, http.StatusOK, toSuccessDelete())
}
