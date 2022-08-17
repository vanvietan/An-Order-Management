package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := ":3000"
	r := chi.NewRouter()

	//Get
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))

	})

	//Post
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {

	})

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)

}
