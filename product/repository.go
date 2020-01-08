package product

import (
	"product-microservice/domain"
)

type Repository interface {
	CreateProduct(domain.Product) interface{}
	ListProducts() interface{}
}
