package tests

import (
	"social_media_app/routes"

	"github.com/gin-gonic/gin"
)

var testToken string

func SetupRouter() *gin.Engine {
	router := gin.Default()
	routes.RegisterRoutes(router)
	return router
}
