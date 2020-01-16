package create_product

import (
	domain "product-microservice/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	CreateProductRequest struct {
		Data struct {
			ProductName        string `json:"product_name"`
			ProductDescription string `json:"product_description"`
			Quantity           int    `json:"qty"`
		}
	}
)

func ValidateRequest(req *CreateProductRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateProductRequest) domain.Product {
	return domain.Product{
		ProductName:        req.Data.ProductName,
		ProductDescription: req.Data.ProductDescription,
		Quantity:           req.Data.Quantity,
	}
}
