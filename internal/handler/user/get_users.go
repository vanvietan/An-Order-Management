package user

import (
	"errors"
	"math"
	"net/http"
	"order-mg/internal/handler/common"
	"strconv"
)

const maxLimit = 100

// GetUsers Get all users
func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	limit, lastID, err := validateAndMap(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}

	users, errH := h.UserSvc.GetUsers(r.Context(), limit, lastID)
	if errH != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	resp := modelToResponseArray(users)
	common.ResponseJson(w, http.StatusOK, toGetUsersResponse(resp))
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
