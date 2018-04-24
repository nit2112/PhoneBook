package routers

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUserRoutes()
	router = SetNumberRoutes()
	return router
}