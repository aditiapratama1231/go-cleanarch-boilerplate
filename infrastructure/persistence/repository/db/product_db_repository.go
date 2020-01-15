package db

import (
	"context"
	"errors"
	"product-microservice/application/infrastructure"
	domain "product-microservice/domain/entities"

	"github.com/jinzhu/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) infrastructure.ProductRepository {
	return &productRepository{
		DB: DB,
	}
}

func (p *productRepository) CreateProduct(ctx context.Context, product domain.Product) error {
	err := p.DB.Create(&product).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) ListProducts(ctx context.Context) interface{} {
	_products := []domain.Product{}

	p.DB.Find(&_products)
	return _products
}

func (p *productRepository) GetProductById(ctx context.Context, id string) (interface{}, error) {
	product := domain.Product{}

	if p.DB.First(&product, id).RecordNotFound() {
		return nil, errors.New("Cannot find product with id " + id)
	}

	return product, nil
}
