package repository

import (
	"order-mg/db"
	"order-mg/internal/model"
)

func GetAllUsers() {

	var listUser []model.Users

	db.GetDatabaseConnection()

}
