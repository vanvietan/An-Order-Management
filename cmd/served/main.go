package main

import (
	"fmt"
	"log"
	"net/http"
	"order-mg/cmd/served/router"
	"order-mg/db"
	userRepo "order-mg/internal/repository/user"
	userSvc "order-mg/internal/service/user"
	"order-mg/internal/util"

	"github.com/go-chi/chi"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func main() {
	port := ":3000"
	r := chi.NewRouter()

	/*
		init repository
		init service
		init controller

		init router
		request -> router (middleware handle) -> GET /order -> controller -> Service -> repository -> DB
		Create DB connection -> repository -> Service -> controller -> Router
	*/

	// Create db connection
	dbConn, err := db.CreateDBConnection()
	if err != nil {
		log.Fatalf("encountered error when create db connection, error: %v", err)
	}

	//sonyflake
	util.Init()

	userRepo := userRepo.New(dbConn)
	userSvc := userSvc.New(userRepo)

	router.New(r, userSvc)

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)

}
