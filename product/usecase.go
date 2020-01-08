package product

import (
	"product-microservice/domain"
)

type Usecase interface {
	CreateProduct(domain.Product) interface{}
	ListProducts() interface{}
}
