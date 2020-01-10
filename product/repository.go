package product

import (
	"context"
	"product-microservice/domain"
)

type Repository interface {
	CreateProduct(context.Context, domain.Product) error
	ListProducts(context.Context) interface{}
	GetProductById(context.Context, string) (interface{}, error)
}
