package router

import (
	"database/sql"
	"fmt"
)

const (
	user     = "postgres"
	password = "admin"
	port     = 5432
	dbname   = "postgres"
	host     = "localhost"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user =%s "+
		"password=%s dbname=%s ssdmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connected!")
	return db
}
