package dao

import (
	"fmt"
	. "github.com/akashg-solulab/golang_challenge/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "Coffee_Machine"
)

// Establish a connection to database
func (m *ProductsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ProductsDAO) FindAll() ([]Coffee_Machine, error) {
	var cm []Coffee_Machine
	err := db.C(COLLECTION).Find(bson.M{}).All(&cm)
	return cm, err
}
