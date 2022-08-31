package user

import (
	"errors"
	"net/http"
	"order-mg/internal/handler/common"
	"order-mg/internal/model"
	"strconv"

	"github.com/go-chi/chi"
)

// GetUserByID  find user by an id
func (h UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := validateCursorAndMap(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	user, err := h.UserSvc.GetUserByID(r.Context(), userID)
	if err != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJson(w, http.StatusOK, toGetAUserResponse(user))

}

func toGetAUserResponse(user model.Users) getAUserResponse {
	return getAUserResponse{
		User: user,
	}
}

func validateCursorAndMap(r *http.Request) (int64, error) {
	ID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, errors.New("id must be a number")
	}
	if ID < 0 {
		return 0, errors.New("invalid id")
	}

	return ID, nil
}
