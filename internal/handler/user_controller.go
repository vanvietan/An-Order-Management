package handler

import (
	"net/http"
	"order-mg/internal/service"

	"github.com/go-chi/chi"
)

var router *chi.Mux

func routers() *chi.Mux {
	router.Get("/users", GetAllUsers)
	router.Get("/users/{id}", GetUser)
	router.Post("/users/", CreateUser)
	router.Put("/users/{id}", UpdateUser)
	router.Delete("/users/{id}", DeleteUser)
	return router
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	service.GetAllUsers()
}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
