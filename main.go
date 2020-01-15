package main

import (
	"fmt"
	"log"
	"os"

	// configs and interfaces
	_repoI "product-microservice/application/infrastructure"
	_db "product-microservice/infrastructure/persistence/repository/db"
	_http "product-microservice/infrastructure/transport/http"

	// actual class
	_repo "product-microservice/infrastructure/persistence/repository/db"

	_createProduct "product-microservice/application/use_case/product/create_product"
	_listProducts "product-microservice/application/use_case/product/list_products"
	_showProduct "product-microservice/application/use_case/product/show_product"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1") // initial route

	db := _db.DBInit() // initial db configuration

	api := _http.NewRequest("")

	// register product
	productRepo := _repo.NewProductRepository(db)
	ProductRoute(v1, api, productRepo)

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

func ProductRoute(route *gin.RouterGroup, req _http.Request, prdRepo _repoI.ProductRepository) {
	cpHandler := _createProduct.NewCreateProductHandler(req, prdRepo)
	lpHandler := _listProducts.NewListProductsHandler(req, prdRepo)
	spHandler := _showProduct.NewShowProductHandler(req, prdRepo)

	v1 := route.Group("/products")
	{
		v1.POST("/", cpHandler.CreateProductHandler)
		v1.GET("/", lpHandler.ListProductsHandler)
		v1.GET("/:id", spHandler.ShowProductHandler)
	}
}
