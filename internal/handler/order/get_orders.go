package order

import (
	"errors"
	"math"
	"net/http"
	"order-mg/internal/handler/common"
	"strconv"
)

const maxLimit = 100

// GetOrders get all orders
func (h OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	limit, lastID, err := validateAndMap(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	orders, errH := h.OrderSvc.GetOrders(r.Context(), limit, lastID)
	if errH != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJson(w, http.StatusOK, toGetOrdersResponse(orders))
}

func validateAndMap(r *http.Request) (int, int64, error) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		return 0, 0, errors.New("limit must be a number")
	}
	cursor, err := strconv.ParseInt(r.URL.Query().Get("cursor"), 10, 64)
	if err != nil {
		return 0, 0, errors.New("cursor must be a number")
	}
	if limit < 1 || limit > maxLimit {
		return 0, 0, errors.New("invalid limit")
	}
	if cursor < 0 || cursor > math.MaxInt64 {
		return 0, 0, errors.New("invalid cursor")
	}

	return limit, cursor, nil
}
