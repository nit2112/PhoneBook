package data

import (
	"time"
	"Phonebook/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NumberRepository struct {
	C *mgo.Collection
}

func (r *NumberRepository) Create(number *models.Number) error {
	obj_id := bson.NewObjectId()
	number.Id = obj_id
	ti := time.Now()
	number.CreatedOn = ti.String()
	err := r.C.Insert(&number)
	return err
}

func (r *NumberRepository) Update(number *models.Number) error {
	err := r.C.Update(bson.M{"_id":number.Id},
		bson.M{"$set":bson.M{
			"name": number.Name,
			"number": number.Mobile,
			"description": number.Description,
			}})
		return err
}

func (r *NumberRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *NumberRepository) GetAll() []models.Number {
	var numbers []models.Number
	iter := r.C.Find(nil).Iter()
	result := models.Number{}
	for iter.Next(&result) {
		numbers = append(numbers, result)
	}
	return numbers
}

func (r *NumberRepository) GetById(id string) (number models.Number, err error){
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&number)
	return
}

func (r *NumberRepository) GetByUser(user string) []models.Number {
	var number []models.Number
	iter := r.C.Find(bson.M{"createdby": user}).Iter()
	result := models.Number{}
	for iter.Next(&result){
		number = append(number, result)
	}
	return number
}