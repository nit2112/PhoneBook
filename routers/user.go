package routers

import (
	"github.com/gorilla/mux"
	"TaskManager/controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("users/register", controllers.Register).Method("POST")
	router.HandleFunc("users/login", controllers.Login).Method("POST")
	return router
}