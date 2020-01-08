package api

import (
	"net/http"
	"product-microservice/domain"
	"product-microservice/product"

	productRequest "product-microservice/product/transport/http/request"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productUsecase product.Usecase
}

func (ph *productHandler) CreateProduct(c *gin.Context) {
	request := productRequest.CreateProductRequest{}

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

	ph.productUsecase.CreateProduct(product)

	c.JSON(http.StatusCreated, gin.H{
		"message": "product successfully created",
	})

	return
}

func (ph *productHandler) ListProducts(c *gin.Context) {
	products := ph.productUsecase.ListProducts()

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
	return
}
