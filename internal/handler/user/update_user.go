package user

import (
	"net/http"
	"order-mg/internal/handler/common"
)

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, errs := checkValidate(r)
	if errs != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid request",
			Description: errs.Error(),
		})
		return
	}
	userID, err := validateCursorAndMap(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid_request",
			Description: err.Error(),
		})
		return
	}
	user, err := h.UserSvc.UpdateUser(r.Context(), reqBody, userID)
	if err != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}

	common.ResponseJson(w, http.StatusOK, toGetAUserResponse(user))

}
