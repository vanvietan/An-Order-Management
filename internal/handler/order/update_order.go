package order

import (
	"errors"
	"github.com/go-chi/chi"
	"math"
	"net/http"
	"order-mg/internal/handler/common"
	"strconv"
)

// UpdateOrder update an order
func (h OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	reqBody, err := checkValidate(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid request",
			Description: err.Error(),
		})
		return
	}
	orderID, errI := validateIDAndMap(r)
	if errI != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid request",
			Description: err.Error(),
		})
		return
	}
	orderUpdated, errH := h.OrderSvc.UpdateOrder(r.Context(), reqBody, orderID)
	if errH != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJson(w, http.StatusOK, toGetAnOrderResponse(orderUpdated))
}

func validateIDAndMap(r *http.Request) (int64, error) {
	ID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, errors.New("ID must be a number")
	}
	if ID <= 0 || ID > math.MaxInt64 {
		return 0, errors.New("invalid ID")
	}
	return ID, nil
}
