package product

type Repository interface {
	CreateProduct() (interface{}, error)
	ListProducts() (interface{}, error)
}
