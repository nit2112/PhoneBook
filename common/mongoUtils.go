package common

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

var session *mgo.Session

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session ,err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs: []string{AppConfig.MongoDBHost},
			Username: AppConfig.MongoDBUser,
			Password: AppConfig.MongoDBPwd,
			Timeout: 60 * time.Second,
			})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

func createDbSession() {
		var err error
		session ,err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs: []string{AppConfig.MongoDBHost},
			Username: AppConfig.MongoDBUser,
			Password: AppConfig.MongoDBPwd,
			Timeout: 60 * time.Second,
			})
		if err != nil {
			log.Fatalf("[createDbSession]: %s\n", err)
		}
}

func addIndexes() {
	var err error
	userIndex := mgo.Index{
		Key:    []string{"email"},
		Unique: true,
		Background: true,
		Sparse: true,
	}
	numberIndex := mgo.Index{
		Key:    []string{"email"},
		Unique: true,
		Background: true,
		Sparse: true,
	}

	session := GetSession().Copy()
	defer session.Close()
	userCol := session.DB(AppConfig.Database).C("users")
	numberCol := session.DB(AppConfig.Database).C("number")
	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	err = numberCol.EnsureIndex(numberIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}


