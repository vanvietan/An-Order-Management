package service

import (
	myUser "order-mg/internal/model"
)

type user = myUser.Users

var userDTO myUser.Users

func GetAllUsers() user {

	// repository.GetAllUsers()

	return userDTO
}
