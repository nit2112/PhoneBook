package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"Phonebook/common"
	"Phonebook/controllers"
)

func SetNumberRoutes(router *mux.Router) *mux.Router {
	numRouter := mux.NewRouter()
	numRouter.HandleFunc("/number",controllers.CreateNumber).Methods("POST")
	numRouter.HandleFunc("/number/{id}",controllers.UpdateNumber).Methods("PUT")
	numRouter.HandleFunc("/number",controllers.GetNumbers).Methods("GET")
	numRouter.HandleFunc("/number/{id}",controllers.GetNumberById).Methods("GET")
	numRouter.HandleFunc("/number/{id}",controllers.DeleteNumber).Methods("DELETE")
	router.PathPrefix("/number").Handler(negroni.New(
			negroni.HandlerFunc(common.Authorize),
			negroni.Wrap(numRouter),
		))

	return router
}