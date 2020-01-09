package repository

import (
	"context"
	"product-microservice/domain"
	"product-microservice/models"
	"product-microservice/product"

	"github.com/jinzhu/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) product.Repository {
	return &productRepository{
		DB: DB,
	}
}

func (p *productRepository) CreateProduct(ctx context.Context, product domain.Product) error {
	_product := models.Product{
		ProductName:        product.ProductName,
		ProductDescription: product.ProductDescription,
		Quantity:           product.Quantity,
	}

	err := p.DB.Create(&_product).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) ListProducts(ctx context.Context) interface{} {
	_product := []models.Product{}

	p.DB.Find(&_product)
	return _product
}
