package models

import (
	// "time"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id             bson.ObjectId       `bson:"_id,omitempty" json:"id"`
		FirstName	   string			   `json:"firstname"`
		LastName	   string			   `json:"lastname"`
		Email		   string			   `json:"email"`
		Password	   string			   `json:"password"`
		HashPassword   []byte			   `json:"hashpassword,omitempty"`
	}

	Number struct {
		Id 				bson.ObjectId     `bson:"_id,omitempty" json:"id"`
		UserId			bson.ObjectId	  `bson:"userid"`
		Name 			string			  `json:"name"`
		Mobile			string			  `json:"number"`
		Description		string 			  `json:"description"`
		CreatedOn 		string			`json:"createdon,omitempty"`
	}
)