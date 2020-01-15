package list_products

import (
	"product-microservice/application/infrastructure"
	domain "product-microservice/domain/entities"
)

type (
	ListTransactionResponse struct {
		infrastructure.BaseResponse
	}
)

func ResponseMapper(domain []domain.Product, message string, success bool) ListTransactionResponse {
	return ListTransactionResponse{
		BaseResponse: infrastructure.BaseResponse{
			Success: success,
			Message: message,
			Data:    domain,
		},
	}
}
