package main

import (
	"github.com/gin-gonic/gin"
	"github.com/MarcoVitoC/pbi-btpns/controllers"
	"github.com/MarcoVitoC/pbi-btpns/database"
)

func main() {
	r := gin.Default()
	database.DatabaseConnection()

	r.POST("/users/register", controllers.Register)
	r.POST("/users/login", controllers.Login)
	r.PUT("/users/:userId", controllers.Update)
	r.DELETE("/users/:userId", controllers.Delete)

	r.Run()
}