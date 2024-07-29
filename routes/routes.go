package routes

import (
	"social_media_app/controllers"
	"social_media_app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser)
	router.POST("/login", controllers.Login)

	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/users/:id", controllers.GetUser)
		authorized.GET("/users", controllers.GetUsers)
		authorized.PUT("/users/:id", controllers.UpdateUser)
		authorized.DELETE("/users/:id", controllers.DeleteUser)
		authorized.POST("/posts", controllers.CreatePost)
		authorized.GET("/posts/:id", controllers.GetPost)
		authorized.GET("/posts", controllers.GetPosts)
		authorized.PUT("/posts/:id", controllers.UpdatePost)
		authorized.DELETE("/posts/:id", controllers.DeletePost)
	}
}
