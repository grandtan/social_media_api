package main

import (
	"log"
	"net/http"
	"social_media_app/database"
	"social_media_app/routes"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
