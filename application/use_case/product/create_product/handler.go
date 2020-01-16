package create_product

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"product-microservice/application/infrastructure"
	"product-microservice/application/misc"

	api "product-microservice/infrastructure/transport/http"

	"github.com/refactory-id/go-core-package/response"
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
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&request); !ok {
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(err.Error(), false))
		return
	}

	prd, err := handler.repository.CreateProduct(ctx, RequestMapper(request))

	if err != nil {
		c.JSON(misc.GetErrorStatusCode(err), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(prd, "Create product success", true))
}
