package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Price       int           `json:"price" bson:"price"`
	Amount      int           `json:"amount" bson:"amount"`
	CreatedTime time.Time     `json:"-" bson:"created_time"`
	UpdatedTime time.Time     `json:"updatedTime" bson:"updated_time"`
}
