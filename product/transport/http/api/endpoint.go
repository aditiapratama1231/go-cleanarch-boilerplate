package api

import (
	"product-microservice/product"

	"github.com/gin-gonic/gin"
)

// product route "/products/..."
func ProductRoute(route *gin.RouterGroup, pu product.Usecase) {
	handler := productHandler{
		productUsecase: pu,
	}

	v1 := route.Group("/products")
	{
		v1.GET("/", handler.ListProducts)
		v1.POST("/", handler.CreateProduct)
	}
}
