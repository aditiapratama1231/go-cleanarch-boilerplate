package create_product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"product-microservice/application/infrastructure"

	domain "product-microservice/domain/entities"

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

func (handler *createProductHandler) CreateProduct(c *gin.Context) {
	request := CreateProductRequest{}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	product := domain.Product{
		ProductName:        request.Data.ProductName,
		ProductDescription: request.Data.ProductDescription,
		Quantity:           request.Data.Quantity,
	}

	err := handler.repository.CreateProduct(c, product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating product" + err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "product successfully created",
	})

	return
}
