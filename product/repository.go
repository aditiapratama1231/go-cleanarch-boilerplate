package product

import (
	"context"
	"product-microservice/domain"
)

type Repository interface {
	CreateProduct(context.Context, domain.Product) interface{}
	ListProducts(context.Context) interface{}
}
