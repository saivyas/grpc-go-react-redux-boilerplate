package databaselayer

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type MongodbHandler struct {
	*mgo.Session
}

func NewMongodbHandler(connection string) (*MongodbHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongodbHandler{
		Session: s,
	}, err
}
func (handler *MongodbHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}


func(handler *MongodbHandler) GetAllBooks() ([]Book,error){
	Books := []Book{}

	return Books,nil
}