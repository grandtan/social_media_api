package main

import (
	"log"
	"social_media_app/database"
	"social_media_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	router := gin.Default()
	routes.RegisterRoutes(router)
	log.Fatal(router.Run(":8080"))
}
