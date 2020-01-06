package product

type Usecase interface {
	CreateProduct() (interface{}, error)
	ListProducts() (interface{}, error)
}
