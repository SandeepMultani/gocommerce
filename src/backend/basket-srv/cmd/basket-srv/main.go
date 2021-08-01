package main

import (
	"github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/core/basket"
	"github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/handlers"
	"github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/repositories"
	"github.com/gin-gonic/gin"
)

func main() {

	basketRedisDb, err := repositories.GetRedisClient()
	if err != nil {
		panic(err)
	}
	basketRepository := repositories.NewBasketRepository(basketRedisDb)
	basketSrv := basket.NewBasketService(basketRepository)
	basketHandler := handlers.NewBasketHandler(basketSrv)

	router := gin.New()

	v1 := router.Group("/v1")
	{
		v1.POST("/basket/", basketHandler.Create)
		v1.GET("/basket/:basketId", basketHandler.Get)
		v1.DELETE("/basket/:basketId", basketHandler.Delete)

		v1.PUT("/basket/:basketId/add", basketHandler.AddItem)
		v1.PUT("/basket/:basketId/remove", basketHandler.RemoveItem)
	}

	err = router.Run(":5002")
	if err != nil {
		panic(err)
	}
}
