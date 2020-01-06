package usecase

import (
	"product-microservice/product"
)

type productUsecase struct {
	productRepo product.Repository
}

func NewProductUsecase(productRepo product.Repository) product.Usecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (p *productUsecase) CreateProduct() (interface{}, error) {
	return nil, nil
}

func (p *productUsecase) ListProducts() (interface{}, error) {
	return nil, nil
}
