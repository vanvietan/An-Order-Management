package user

import (
	"errors"
	"math"
	"net/http"
	"order-mg/internal/handler/common"
	"strconv"

	"github.com/go-chi/chi"
)

// DeleteUser delete a user
func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID, err := validateIDAndMap(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	errD := h.UserSvc.DeleteUser(r.Context(), userID)
	if errD != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJson(w, http.StatusOK, toSuccessDelete())
}
func toSuccessDelete() deleteUserResponse {
	return deleteUserResponse{
		Message: "Deleted User",
	}
}
func validateIDAndMap(r *http.Request) (int64, error) {
	ID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, errors.New("id must be a number")
	}
	if ID <= 0 || ID > math.MaxInt64 {
		return 0, errors.New("invalid id")
	}

	return ID, nil
}
