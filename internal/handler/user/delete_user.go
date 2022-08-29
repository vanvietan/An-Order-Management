package user

import (
	"net/http"
	"order-mg/internal/handler/common"
)

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID, err := validateCursorAndMap(r)
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
