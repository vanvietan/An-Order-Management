package user

import (
	"errors"
	"net/http"
	"order-mg/internal/handler/common"
	"strconv"

	"github.com/go-chi/chi"
)

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID, err := validateIDAndMap(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	isSuccess, err := h.UserSvc.DeleteUser(r.Context(), userID)
	if err != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJson(w, http.StatusOK, toSuccessDelete(isSuccess))
}
func toSuccessDelete(isSuccess bool) deleteUserResponse {
	return deleteUserResponse{
		Status:  isSuccess,
		Message: "Deleted User",
	}
}
func validateIDAndMap(r *http.Request) (int64, error) {
	cursor, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, errors.New("id must be a number")
	}
	if cursor < 0 {
		return 0, errors.New("invalid id")
	}

	return cursor, nil
}
