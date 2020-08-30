package repository

import (
	"github.com/dacharat/gomongo/src/model"
	"github.com/globalsign/mgo"
)

type ProductRepository interface {
	GetAllProduct() ([]model.Product, error)
}

type ProductRepositoryImpl struct {
	Connection *mgo.Session
}

const (
	DbName     = "Shop"
	collection = "product"
)

func (p *ProductRepositoryImpl) GetAllProduct() ([]model.Product, error) {
	var products []model.Product
	err := p.Connection.DB(DbName).C(collection).Find(nil).All(&products)
	return products, err

}
