package api

import (
	"log"
	"net/http"

	"github.com/dacharat/gomongo/src/model"
	"github.com/dacharat/gomongo/src/repository"
	"github.com/gin-gonic/gin"
)

type ProductApi struct {
	ProductRepository repository.ProductRepository
}

func (api ProductApi) ProductListHandler(context *gin.Context) {
	var products model.Products
	allProduct, err := api.ProductRepository.GetAllProduct()
	if err != nil {
		log.Println("error productListHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	products.Products = allProduct

	context.JSON(http.StatusOK, products)
}
