package tests

import (
	"social_media_app/routes"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    routes.RegisterRoutes(router)
    return router
}
