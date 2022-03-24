package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	_c := controllers.UserControllers{}

	r.GET("/api/1/users", _c.GetUsers)
	r.GET("/api/1/users/:id", _c.GetUser)
	r.POST("/api/1/users", _c.AddUser)
	r.PUT("/api/1/users/:id", _c.UpdateUser)
	r.DELETE("/api/1/users/:id", _c.DeleteUser)
}
