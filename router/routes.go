package router

import (
	"github.com/gin-gonic/gin"
	"github.com/MarcoVitoC/pbi-btpns/controllers"
	"github.com/MarcoVitoC/pbi-btpns/database"
	"github.com/MarcoVitoC/pbi-btpns/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	database.DatabaseConnection()

	r.POST("/users/register", controllers.Register)
	r.POST("/users/login", controllers.Login)
	r.GET("/users/validate", middlewares.Auth, controllers.Validate)
	r.PUT("/users/:userId", controllers.Update)
	r.DELETE("/users/:userId", controllers.Delete)

	return r
}