package router

import (
	"github.com/gin-gonic/gin"
	"github.com/MarcoVitoC/pbi-btpns/controllers"
	"github.com/MarcoVitoC/pbi-btpns/middlewares"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", controllers.Register)
		userRoutes.POST("/login", controllers.Login)
		userRoutes.PUT("/:userId", controllers.Update)
		userRoutes.DELETE("/:userId", controllers.Delete)
	}

	photoRoutes := r.Group("/photos")
	{
		photoRoutes.POST("/", middlewares.Auth, controllers.UploadPhoto)
		photoRoutes.GET("/", controllers.GetPhoto)
		photoRoutes.PUT("/:photoId", middlewares.Auth, controllers.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", middlewares.Auth, controllers.DeletePhoto)
	}

	return r
}