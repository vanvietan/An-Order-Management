package user

import (
	"net/http"
	"order-mg/internal/handler/common"
)

// GetUserByID  find user by an id
func (h UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := validateIDAndMap(r)
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
