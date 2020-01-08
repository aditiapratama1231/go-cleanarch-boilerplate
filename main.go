package main

import (
	"fmt"
	"log"
	"os"

	"product-microservice/config"
	_productRepo "product-microservice/product/repository"
	_productApi "product-microservice/product/transport/http/api"
	_productUsecase "product-microservice/product/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1") // initial route

	db := config.DBInit() // initial db configuration

	// register repo and usecase
	productRepo := _productRepo.NewProductRepository(db)             // initial product repository
	productUsecase := _productUsecase.NewProductUsecase(productRepo) // initial product usecase

	// register routing
	_productApi.ProductRoute(v1, productUsecase)

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal(fmt.Sprintf("PORT must be set [%s]", port))
	}

	router.Run(":" + port)
}
