package main

import (
	"gin-mongo-api/configs"
	routes "gin-mongo-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// run database
	configs.ConnectDB()
	defer configs.DisconnetDB()

	routes.RegisterUserRoutes(router)

	router.Run(":8000")
}
