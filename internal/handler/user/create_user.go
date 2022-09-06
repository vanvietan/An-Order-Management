package user

import (
	"encoding/json"
	"net/http"
	"order-mg/internal/handler/common"
	"order-mg/internal/model"
)

// CreateUser create a user
func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := checkValidate(r)
	if err != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid request",
			Description: err.Error(),
		})
		return
	}

	userCreated, err := h.UserSvc.CreateUser(r.Context(), reqBody)
	if err != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	common.ResponseJson(w, http.StatusOK, toGetAUserResponse(userCreated))
}

func checkValidate(r *http.Request) (model.Users, error) {
	var input CreateUserInput
	_ = json.NewDecoder(r.Body).Decode(&input)
	svcInput, err := input.validateAndMap()
	if err != nil {
		return model.Users{}, err
	}
	return svcInput, nil
}
