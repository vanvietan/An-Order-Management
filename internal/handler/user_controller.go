package handler

import (
	"net/http"
	"order-mg/internal/service"

	"github.com/go-chi/chi"
)

var router *chi.Mux

func routers() *chi.Mux {
	router.Get("/users", GetAllUsers)
	router.Get("/users/{id}", GetUserById)
	router.Post("/users/", CreateUser)
	router.Put("/users/{id}", UpdateUser)
	router.Delete("/users/{id}", DeleteUser)
	return router
}

// Get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	service.GetAllUsers()
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
