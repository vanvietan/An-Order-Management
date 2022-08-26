package user

import (
	"encoding/json"
	"errors"
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
	var user model.Users
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Password == "" {
		return model.Users{}, errors.New("invalid password")
	}
	if user.Username == "" {
		return model.Users{}, errors.New("invalid username")
	}

	return user, nil
}
