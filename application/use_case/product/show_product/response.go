package show_product

import (
	domain "product-microservice/domain/entities"
	"time"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowProductResponse struct {
		base.BaseResponse
		Data ShowProductResponseData `json:"data"`
	}

	ShowProductResponseData struct {
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

func SetResponse(domain domain.Product, message string, success bool) ShowProductResponse {
	return ShowProductResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: ResponseMapper(domain),
	}
}

func ResponseMapper(domain domain.Product) ShowProductResponseData {
	return ShowProductResponseData{
		ID:                 domain.Model.ID,
		ProductName:        domain.ProductName,
		ProductDescription: domain.ProductDescription,
		Quantity:           domain.Quantity,
		CreatedAt:          domain.Model.CreatedAt,
		UpdatedAt:          domain.Model.UpdatedAt,
		DeletedAt:          domain.Model.DeletedAt,
	}
}
