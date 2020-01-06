package main

import (
	"fmt"
	"log"
	"os"

	"product-microservice/config"
	"product-microservice/product/repository"
	"product-microservice/product/transport/http/api"
	"product-microservice/product/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	db := config.GetDB()

	_productRepo := repository.NewProductRepository(db)
	_productUsecase := usecase.NewProductUsecase(_productRepo)
	api.NewProductHandler(v1, _productUsecase)

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
