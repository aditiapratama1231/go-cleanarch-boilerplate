package show_product

import (
	"context"
	"net/http"
	"product-microservice/application/infrastructure"
	"product-microservice/application/misc"

	"github.com/gin-gonic/gin"
	api "product-microservice/infrastructure/transport/http"

	"github.com/refactory-id/go-core-package/response"
)

type showProductHandler struct {
	request       api.Request
	prdRepository infrastructure.ProductRepository
}

func NewShowProductHandler(request api.Request, prdRepo infrastructure.ProductRepository) showProductHandler {
	return showProductHandler{
		request:       request,
		prdRepository: prdRepo,
	}
}

func (handler *showProductHandler) ShowProductHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	product, err := handler.prdRepository.GetProductById(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(err), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(product, "Show product success", true))
}
