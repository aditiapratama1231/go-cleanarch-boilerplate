package product

import (
	"context"
	"product-microservice/domain"
)

type Usecase interface {
	CreateProduct(context.Context, domain.Product) error
	ListProducts(context.Context) (interface{}, error)
}
