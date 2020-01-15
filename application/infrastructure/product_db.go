package infrastructure

import (
	"context"
	domain "product-microservice/domain/entities"
)

type ProductRepository interface {
	CreateProduct(context.Context, domain.Product) error
	ListProducts(context.Context) []domain.Product
	GetProductById(context.Context, string) (domain.Product, error)
}
