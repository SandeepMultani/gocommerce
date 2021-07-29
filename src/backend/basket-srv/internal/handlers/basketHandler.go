package handlers

import (
	basket "github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/core/basket"
	"github.com/gin-gonic/gin"
)

type BasketHandler interface {
	Get(*gin.Context)
	Create(*gin.Context)
	Delete(*gin.Context)
	AddItem(*gin.Context)
	RemoveItem(*gin.Context)
}

type basketHandler struct {
	basketSrv basket.BasketService
}

type httpResponse struct {
	Message string
	Code    int
}

var _ BasketHandler = &basketHandler{}

func NewBasketHandler(basketSrv basket.BasketService) BasketHandler {
	return &basketHandler{
		basketSrv: basketSrv,
	}
}

func (hdl *basketHandler) Get(c *gin.Context) {
	basketId := c.Param("basketId")
	bas, err := hdl.basketSrv.Get(basketId)
	if err != nil {
		c.JSON(404, httpResponse{
			Message: err.Error(),
			Code:    404,
		})
		return
	}
	c.JSON(200, &bas)
}

func (hdl *basketHandler) Create(c *gin.Context) {
	basketId := c.Param("basketId")
	bas, err := hdl.basketSrv.Create(basketId)
	if err != nil {
		c.JSON(404, httpResponse{
			Message: err.Error(),
			Code:    404,
		})
		return
	}
	c.JSON(200, &bas)
}

func (hdl *basketHandler) Delete(c *gin.Context) {
	basketId := c.Param("basketId")
	err := hdl.basketSrv.Delete(basketId)
	if err != nil {
		c.JSON(404, httpResponse{
			Message: err.Error(),
			Code:    404,
		})
		return
	}
	c.JSON(200, httpResponse{
		Message: "success",
		Code:    200,
	})
}

func (hdl *basketHandler) AddItem(c *gin.Context) {
	basketId := c.Param("basketId")
	productId := c.Param("productId")
	bas, err := hdl.basketSrv.AddItem(basketId, productId)
	if err != nil {
		c.JSON(404, httpResponse{
			Message: err.Error(),
			Code:    404,
		})
		return
	}
	c.JSON(200, bas)
}

func (hdl *basketHandler) RemoveItem(c *gin.Context) {
	basketId := c.Param("basketId")
	productId := c.Param("productId")
	bas, err := hdl.basketSrv.RemoveItem(basketId, productId)
	if err != nil {
		c.JSON(404, httpResponse{
			Message: err.Error(),
			Code:    404,
		})
		return
	}
	c.JSON(200, bas)
}
