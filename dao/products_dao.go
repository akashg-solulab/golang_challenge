package dao

import (
	"fmt"
	"github.com/akashg-solulab/golang_challenge/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database
fmt.printf("db")

const (
	COLLECTION = "products"
)

// Establish a connection to database
func (p *ProductsDAO) Connect() {
	session, err := mgo.Dial(p.Server)
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


// Find a movie by its id
func (m *ProductsDAO) FindById(id string) (Products, error) {
	var products Products
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&products)
	return products, err
}

// Insert a movie into database
func (m *ProductsDAO) Insert(products Products) error {
	err := db.C(COLLECTION).Insert(&products)
	return err
}

// Delete an existing movie
func (m *ProductsDAO) Delete(products Products) error {
	err := db.C(COLLECTION).Remove(&products)
	return err
}

// Update an existing movie
func (m *ProductsDAO) Update(products Products) error {
	err := db.C(COLLECTION).UpdateId(products.ID, &products)
	return err
}
