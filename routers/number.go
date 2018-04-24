package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"TaskManager/common"
	"TaskManager/controllers"
)

func SetNumberRoutes(router *mux.Router) *mux.Router {
	numRouter := mux.NewRouter()
	numRouter.HandleFunc("/number",controllers.CreateNumber).Method("POST")
	numRouter.HandleFunc("/number/{id}",controllers.UpdateNumber).Method("PUT")
	numRouter.HandleFunc("/number",controllers.GetNumbers).Method("GET")
	numRouter.HandleFunc("/number/{id}",controllers.GetNumber).Method("GET")
	numRouter.HandleFunc("/number/{id}",controllers.DeleteNumber).Method("DELETE")
	router.PathPrefix("/number").Handler(negroni.New(
			negroni.HandlerFunc(common.Authrize),
			negroni.Wrap(numRouter),
		))

	return router
}