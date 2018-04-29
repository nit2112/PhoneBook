package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"Phonebook/common"
	"Phonebook/routers"
)

func main(){
	common.StartUp()
	router := routers.InitRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr: common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listing...")
	server.ListenAndServe()
}