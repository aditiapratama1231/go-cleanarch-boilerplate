package repository

import (
	"github.com/jinzhu/gorm"
	"product-microservice/domain"
	"product-microservice/models"
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

func (p *productRepository) CreateProduct(product domain.Product) interface{} {
	_product := models.Product{
		ProductName:        product.ProductName,
		ProductDescription: product.ProductDescription,
		Quantity:           product.Quantity,
	}

	p.DB.Create(&_product)
	return nil
}

func (p *productRepository) ListProducts() interface{} {
	_product := []models.Product{}

	p.DB.Find(&_product)
	return _product
}
