package main

import (
	"github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/internal/core/product"
	"github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/internal/handlers"
	"github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/internal/repositories"
	"github.com/gin-gonic/gin"
)

func main() {

	productRepository := repositories.NewProductRepository()
	productSrv := product.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productSrv)

	router := gin.New()

	v1 := router.Group("/v1")
	{
		v1.GET("/product", productHandler.Get)
		v1.POST("/product", productHandler.Create)
		v1.GET("/product/sku", productHandler.GetBySku)
		v1.GET("/product/:id", productHandler.GetById)
	}

	err := router.Run(":5001")
	if err != nil {
		panic(err)
	}
}
