package create_product

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"product-microservice/application/infrastructure"
	"product-microservice/application/misc"

	api "product-microservice/infrastructure/transport/http"
)

type createProductHandler struct {
	request    api.Request
	repository infrastructure.ProductRepository
}

func NewCreateProductHandler(request api.Request, prdRepo infrastructure.ProductRepository) createProductHandler {
	return createProductHandler{
		request:    request,
		repository: prdRepo,
	}
}

func (handler *createProductHandler) CreateProductHandler(c *gin.Context) {
	request := CreateProductRequest{}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMapper(nil, err.Error(), false))
		return
	}

	product := RequestMapper(request)

	err := handler.repository.CreateProduct(ctx, product)

	if err != nil {
		c.JSON(misc.GetErrorStatusCode(err), ResponseMapper(nil, err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, ResponseMapper(&product, "Create product success", true))
	return
}
