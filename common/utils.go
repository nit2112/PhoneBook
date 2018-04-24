package common

import (
	"encoding/json"
	"log"
	"os"
	"net/http"
)

type configuration struct {
	Server, MongoDBHost, MongoDBUser, MongoDBPwd, Database string
}

type(
	appError struct{
		Error       string   `json:"error"`
		Message     string   `json:"message"`
		HttpStatus  int	     `json:"status"`
	}
	errorResource struct{
		Data appError `json:"data"`
	}
)

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int){
	errObj := appError{
		Error:   handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}
	log.Printf("AppError]:%s\n",handlerError)
	w.Header().Set("Content-Type","application/json;charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil{
		w.Write(j)
	}
}

var AppConfig configuration

func initConfig(){
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadconfig]:%s\n",err)
		}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]:%s\n", err)
	}	
}