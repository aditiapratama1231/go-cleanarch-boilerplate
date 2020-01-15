package show_product

import (
	"product-microservice/application/infrastructure"
	domain "product-microservice/domain/entities"
)

type (
	ShowTransactionResponse struct {
		infrastructure.BaseResponse
	}
)

func ResponseMapper(domain *domain.Product, message string, success bool) ShowTransactionResponse {
	return ShowTransactionResponse{
		BaseResponse: infrastructure.BaseResponse{
			Success: success,
			Message: message,
			Data:    domain,
		},
	}
}
