package controllers

import (
	"fmt"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"gin-mongo-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllers struct {
}

// Get All User
func (c UserControllers) GetUsers(ctx *gin.Context) {
	result, err := services.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
}

// Get User
func (c UserControllers) GetUser(ctx *gin.Context) {
	fmt.Println("get user")
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	user, err := services.GetUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
}

// Create new User
func (c UserControllers) AddUser(ctx *gin.Context) {
	var user models.User
	// Check binding data from ctx
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	// Create User
	result, err := services.AddUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	ctx.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}

// Update User
func (c UserControllers) UpdateUser(ctx *gin.Context) {
	var user *models.User
	id, _ := primitive.ObjectIDFromHex(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	result, err := services.UpdateUser(id, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	ctx.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
}

// Delete User
func (c UserControllers) DeleteUser(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	err = services.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	ctx.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": id}})
}
