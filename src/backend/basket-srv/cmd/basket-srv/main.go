package main

import (
	"github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/core/basket"
	"github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/handlers"
	"github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/repositories"
	"github.com/gin-gonic/gin"
)

func main() {

	basketRepository := repositories.NewBasketRepository()
	basketSrv := basket.NewBasketService(basketRepository)
	basketHandler := handlers.NewBasketHandler(basketSrv)

	router := gin.New()

	v1basket := router.Group("/v1/basket")
	{
		v1basket.GET("/:basketId", basketHandler.Get)
		v1basket.POST("/:basketId", basketHandler.Create)
		v1basket.DELETE("/:basketId", basketHandler.Delete)

		v1basket.POST("/:basketId/product/:productId", basketHandler.AddItem)
		v1basket.DELETE("/:basketId/product/:productId", basketHandler.RemoveItem)
	}

	router.Run(":5002")
}
