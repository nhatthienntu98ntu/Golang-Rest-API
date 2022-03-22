package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nhatthienntu98ntu/Golang-Rest-API/controller"
)

func RegisterUsersRouter(r *mux.Router) {
	_r := r.PathPrefix("/api/1/users").Subrouter()
	_c := controller.UsersController{}

	_r.HandleFunc("", _c.GetAll).Methods(http.MethodGet)
	_r.HandleFunc("/{id}", _c.GetUserById).Methods(http.MethodGet)
	_r.HandleFunc("", _c.AddUser).Methods(http.MethodPost)
	_r.HandleFunc("/{id}", _c.UpdateUser).Methods(http.MethodPut)
	_r.HandleFunc("/{id}", _c.DeleteUser).Methods(http.MethodDelete)
}
