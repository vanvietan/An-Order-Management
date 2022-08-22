package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

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
	//Get
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to An-MG"))

	})

	//Post
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {

	})

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)

}
