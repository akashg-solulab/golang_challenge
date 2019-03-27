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
	COLLECTION = "products"
)

// Establish a connection to database
func (m *ProductsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ProductsDAO) FindAll() ([]Products, error) {
	var products []Products
	err := db.C(COLLECTION).Find(bson.M{}).All(&products)
	return products, err
}
