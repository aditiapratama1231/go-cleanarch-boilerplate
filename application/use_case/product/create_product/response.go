package create_product

import (
	domain "product-microservice/domain/entities"
	"time"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	CreateProductResponse struct {
		base.BaseResponse
		Data CreateProductResponseData `json:"data"`
	}

	CreateProductResponseData struct {
		ID                 uint64     `json:"id"`
		ProductName        string     `json:"product_name"`
		ProductDescription string     `json:"product_description"`
		Quantity           int        `json:"quantity"`
		CreatedAt          time.Time  `json:"created_at"`
		UpdatedAt          time.Time  `json:"updated_at"`
		DeletedAt          *time.Time `json:"deleted_at"`
	}
)

func SetMessage(message string, success bool) base.BaseResponse {
	return base.BaseResponse{
		Message: message,
		Success: success,
	}
}

func SetResponse(domain domain.Product, message string, success bool) CreateProductResponse {
	return CreateProductResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: ResponseMapper(domain),
	}
}

func ResponseMapper(domain domain.Product) CreateProductResponseData {
	return CreateProductResponseData{
		ID:                 domain.Model.ID,
		ProductName:        domain.ProductName,
		ProductDescription: domain.ProductDescription,
		Quantity:           domain.Quantity,
		CreatedAt:          domain.Model.CreatedAt,
		UpdatedAt:          domain.Model.UpdatedAt,
		DeletedAt:          domain.Model.DeletedAt,
	}
}
