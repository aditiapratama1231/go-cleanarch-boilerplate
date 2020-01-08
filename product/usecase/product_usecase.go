package usecase

import (
	"product-microservice/product"

	"product-microservice/domain"
)

type productUsecase struct {
	productRepo product.Repository
}

func NewProductUsecase(productRepo product.Repository) product.Usecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (p *productUsecase) CreateProduct(product domain.Product) interface{} {
	p.productRepo.CreateProduct(product)
	return nil
}

func (p *productUsecase) ListProducts() interface{} {
	return p.productRepo.ListProducts()
}
