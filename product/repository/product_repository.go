package repository

import (
	"github.com/jinzhu/gorm"
	"product-microservice/product"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) product.Repository {
	return &productRepository{
		DB: DB,
	}
}

func (p *productRepository) CreateProduct() (interface{}, error) {
	return nil, nil
}

func (p *productRepository) ListProducts() (interface{}, error) {
	return nil, nil
}
