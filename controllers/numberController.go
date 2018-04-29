package controllers

import (
	"encoding/json"
	// "log"
	"net/http"
	"Phonebook/common"
	"github.com/gorilla/mux"
	"Phonebook/data"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateNumber(w http.ResponseWriter, r *http.Request){
	var dataResource NumberResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Number Data",
			500,
			)
		return
	}
	number := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("number")
	repo := &data.NumberRepository{c}

	repo.Create(number)
	if j, err := json.Marshal(NumberResource{Data: *number}); err !=nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

func GetNumbers(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context .DbCollection("numbers")
	repo := &data.NumberRepository{c}
	number := repo.GetAll()
	j, err := json.Marshal(NumbersResource{Data: number})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type","application/json")
	w.Write(j)
}

func GetNumberById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context .DbCollection("number")
	repo := &data.NumberRepository{c}
	number,err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}
	}
	if j, err := json.Marshal(number); err != nil {
		common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		} else {
    w.Header().Set("content-Type","application/json")
	w.WriteHeader(http.StatusOK)	
	w.Write(j)
}
}

func GetNumbersByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("numbers")
	repo := &data.NumberRepository{c}
	numbers := repo.GetByUser(user)
	j, err := json.Marshal(NumbersResource{Data: numbers})
	if err != nil {
		common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
	}
	w.WriteHeader(http.StatusOK)	
    w.Header().Set("content-Type","application/json")
	w.Write(j)
}

func UpdateNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource NumberResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
	}

	number := &dataResource.Data
	number.Id = id
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("numbers")
	repo := &data.NumberRepository{c}
	if err := repo.Update(number); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.Close()
	c := context .DbCollection("numbers")
	repo := &data.NumberRepository{c}
	err := repo.Delete(id)
	if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
	}
	w.WriteHeader(http.StatusNoContent)	
}






