package create_product

import (
	"product-microservice/application/infrastructure"
	domain "product-microservice/domain/entities"
)

type (
	CreateTransactionResponse struct {
		infrastructure.BaseResponse
	}
)

func ResponseMapper(domain *domain.Product, message string, success bool) CreateTransactionResponse {
	return CreateTransactionResponse{
		BaseResponse: infrastructure.BaseResponse{
			Success: success,
			Message: message,
			Data:    domain,
		},
	}
}
