package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nhatthienntu98ntu/Golang-Rest-API/model"
	"github.com/nhatthienntu98ntu/Golang-Rest-API/service"
	"github.com/nhatthienntu98ntu/Golang-Rest-API/ultils"
)

type UsersController struct {
}

func (u UsersController) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := service.GetAll()

	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
	} else {
		ultils.Return(w, true, http.StatusOK, nil, result)
	}
}

func (u UsersController) GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
		return
	}
	user, err := service.GetUserById(id)
	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
	} else {
		ultils.Return(w, true, http.StatusOK, nil, user)
	}
}

func (u UsersController) AddUser(w http.ResponseWriter, r *http.Request) {
	// Get data from body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
		return
	}
	keyValue := make(map[string]string)
	json.Unmarshal(body, &keyValue)
	// Get User
	user := model.UserInput{
		Username:    keyValue["username"],
		Phone:       keyValue["phone"],
		DateOfBirth: keyValue["dateOfBirth"],
	}
	fmt.Println(user)
	// Create User
	result, err := service.AddUser(user)
	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
	} else {
		ultils.Return(w, true, http.StatusOK, nil, result)
	}
}

func (u UsersController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get id from params
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
		return
	}
	// Get user
	// Get data from body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
		return
	}
	keyValue := make(map[string]string)
	json.Unmarshal(body, &keyValue)
	user := model.UserInput{
		Username:    keyValue["username"],
		Phone:       keyValue["phone"],
		DateOfBirth: keyValue["dateOfBirth"],
	}
	// Update User
	result, err := service.UpdateUser(id, user)
	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
	} else {
		ultils.Return(w, true, http.StatusOK, nil, result)
	}
}

func (u UsersController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
		return
	}
	result, err := service.DeleteUser(id)
	if err != nil {
		ultils.Return(w, true, http.StatusNotFound, err, nil)
	} else {
		ultils.Return(w, true, http.StatusOK, nil, result)
	}
}
