package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"order-mg/internal/handler/common"
	"order-mg/internal/model"
)

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, errs := checkValidateUpdate(r)
	if errs != nil {
		common.ResponseJson(w, http.StatusBadRequest, common.CommonErrorResponse{
			Code:        "invalid request",
			Description: errs.Error(),
		})
		return
	}
	userID, err := validateIDAndMap(r)
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
func checkValidateUpdate(r *http.Request) (model.Users, error) {
	var user model.Users
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Password == "" || len(user.Password) > 14 {
		return model.Users{}, errors.New("password is invalid")
	}
	if user.PhoneNumber == "" || len(user.PhoneNumber) > 11 || len(user.PhoneNumber) < 10 {
		return model.Users{}, errors.New("phone number is invalid")
	}
	if user.Address == "" || len(user.Address) > 120 {
		return model.Users{}, errors.New("address is invalid")
	}
	if user.Age <= 0 || user.Age > 100 {
		return model.Users{}, errors.New("user age is invalid")
	}
	if user.Role != "USER" {
		return model.Users{}, errors.New("user role is invalid")
	}

	return user, nil
}
