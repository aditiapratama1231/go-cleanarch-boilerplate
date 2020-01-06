package api

import (
	"product-microservice/product"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productUsecase product.Usecase
}

func (ph *productHandler) CreateProduct(c *gin.Context) {
	ph.productUsecase.CreateProduct()
	return
}

func (ph *productHandler) ListProducts(c *gin.Context) {
	ph.productUsecase.ListProducts()
	return
}
