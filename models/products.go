package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Products struct {
	ID                         bson.ObjectId `bson:"_id" json:"id"`
	Product                    string        `bson:"product" json:"product"`
	ProductType                string        `bson:"product_type" json:"product_type"`
	SKUCode                    string        `bson:"sku_code" json:"sku_code"`
	WaterLineCompatible        bool          `bson:"water_line_compatible" json:"water_line_compatible"`
	Model                      string        `bson:"model" json:"model"`
	CoffeeFlavor               string        `bson:"coffee_flavor" json:"coffee_flavor"`
	PackSize                   int           `bson:"pack_size" json:"pack_size"`
}
