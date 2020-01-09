package product

import (
	"context"
	"product-microservice/domain"
)

type Usecase interface {
	CreateProduct(context.Context, domain.Product) interface{}
	ListProducts(context.Context) (interface{}, error)
}
