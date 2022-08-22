package repository

import (
	"order-mg/db"
	"order-mg/internal/model"
)

func GetAllUsers() {

	listUsers := []model.Users{}

	dbConn, _ := db.GetDatabaseConnection()

	dbConn.Find(&listUsers)
	// return listUsers

}
