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

	v1product := router.Group("/v1/product")
	{
		v1product.GET("/sku", productHandler.GetBySku)
		v1product.GET("/:id", productHandler.Get)
	}

	err := router.Run(":5001")
	if err != nil {
		panic(err)
	}
}
