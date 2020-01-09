package api

import (
	"context"
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

	err := ph.productUsecase.CreateProduct(c, product)

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

func (ph *productHandler) ListProducts(c *gin.Context) {
	ctx := c.Request.Context()

	if ctx == nil {
		ctx = context.Background()
	}

	products, err := ph.productUsecase.ListProducts(ctx)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": "request timeout",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
	return
}
