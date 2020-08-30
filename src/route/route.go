package route

import (
	"github.com/dacharat/gomongo/src/api"
	"github.com/dacharat/gomongo/src/repository"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func NewRouteProduct(route *gin.Engine, connection *mgo.Session) {
	productRepository := repository.ProductRepositoryImpl{Connection: connection}
	productApi := api.ProductApi{ProductRepository: &productRepository}

	route.GET("api/v1/product", productApi.ProductListHandler)
}
