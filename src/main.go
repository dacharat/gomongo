package main

import (
	"log"

	"github.com/dacharat/gomongo/src/route"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

const (
	// mongoUrl = "mongodb://localhost:27017"
	mongoUrl = "gomongodb:27017"
	port     = ":8080"
)

func main() {
	connection, err := mgo.Dial(mongoUrl)
	if err != nil {
		log.Panic("Cannot connect to Database", err.Error())
	}
	router := gin.Default()
	route.NewRouteProduct(router, connection)
	router.Run(port)
}
