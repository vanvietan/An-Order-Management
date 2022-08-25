package user

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"order-mg/internal/model"
)

// CreateUser create a user
func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := checkValidate(r)

}

func checkValidate(r *http.Request) (model.Users, error) {
	userCheck, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return model.Users{}, errors.New("user invalid field")
	}
	bodyUser := bytes.NewReader(userCheck)

	bodyUser.
	
}
