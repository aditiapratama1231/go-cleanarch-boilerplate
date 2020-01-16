package list_products

import (
	"context"
	"net/http"
	"product-microservice/application/infrastructure"

	api "product-microservice/infrastructure/transport/http"

	"github.com/gin-gonic/gin"
)

type listProductsHandler struct {
	request       api.Request
	prdRepository infrastructure.ProductRepository
}

func NewListProductsHandler(request api.Request, prdRepo infrastructure.ProductRepository) listProductsHandler {
	return listProductsHandler{
		request:       request,
		prdRepository: prdRepo,
	}
}

func (handler *listProductsHandler) ListProductsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	products := handler.prdRepository.ListProducts(ctx)

	c.JSON(http.StatusOK, SetResponse(products, "List Products Success", true))
	return
}
