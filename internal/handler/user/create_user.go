package user

import (
	"encoding/json"
	"errors"
	"fmt"
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
	fmt.Printf("Handler: %+v", userCreated)
	if err != nil {
		common.ResponseJson(w, http.StatusInternalServerError, common.InternalCommonErrorResponse)
		return
	}
	common.ResponseJson(w, http.StatusOK, toGetAUserResponse(userCreated))
}

func checkValidate(r *http.Request) (model.Users, error) {
	var user model.Users
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "" || len(user.Username) > 14 {
		return model.Users{}, errors.New("invalid username")
	}
	if user.Password == "" || len(user.Password) > 14 {
		return model.Users{}, errors.New("password is invalid")
	}
	if user.PhoneNumber == "" || len(user.PhoneNumber) > 11 || len(user.PhoneNumber) < 10 {
		return model.Users{}, errors.New("phone number is invalid")
	}
	if user.Address == "" || len(user.Address) > 120 {
		return model.Users{}, errors.New("address is invalid")
	}
	if user.Age <= 0 || user.Age > 99 {
		return model.Users{}, errors.New("user age is invalid")
	}
	if user.Role != "USER" {
		return model.Users{}, errors.New("user role is invalid")
	}

	return user, nil
}
