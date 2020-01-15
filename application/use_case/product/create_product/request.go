package create_product

import (
	domain "product-microservice/domain/entities"
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

func RequestMapper(req CreateProductRequest) domain.Product {
	return domain.Product{
		ProductName:        req.Data.ProductName,
		ProductDescription: req.Data.ProductDescription,
		Quantity:           req.Data.Quantity,
	}
}
