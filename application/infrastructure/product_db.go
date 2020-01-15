package infrastructure

import (
	"context"
	domain "product-microservice/domain/entities"
)

type ProductRepository interface {
	CreateProduct(context.Context, domain.Product) error
	ListProducts(context.Context) interface{}
	GetProductById(context.Context, string) (interface{}, error)
}
