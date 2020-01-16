package list_products

import (
	"time"

	base "github.com/refactory-id/go-core-package/response"
	domain "product-microservice/domain/entities"
)

type (
	ListProductsResponse struct {
		base.BaseResponse
		Data []ListProductsResponseData `json:"data"`
	}

	ListProductsResponseData struct {
		ID                 uint64     `json:"id"`
		ProductName        string     `json:"product_name"`
		ProductDescription string     `json:"product_description"`
		Quantity           int        `json:"quantity"`
		CreatedAt          time.Time  `json:"created_at"`
		UpdatedAt          time.Time  `json:"updated_at"`
		DeletedAt          *time.Time `json:"deleted_at"`
	}
)

func (res *ListProductsResponse) AddDomain(products []domain.Product) {
	response := ListProductsResponseData{}

	for _, prd := range products {
		response.ID = prd.Model.ID
		response.ProductName = prd.ProductName
		response.ProductDescription = prd.ProductDescription
		response.CreatedAt = prd.CreatedAt
		response.UpdatedAt = prd.UpdatedAt
		response.DeletedAt = prd.DeletedAt

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []domain.Product, message string, success bool) ListProductsResponse {
	return ListProductsResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: ResponseMapper(domain),
	}
}

func ResponseMapper(domain []domain.Product) []ListProductsResponseData {
	response := ListProductsResponse{}

	response.AddDomain(domain)
	return response.Data
}
