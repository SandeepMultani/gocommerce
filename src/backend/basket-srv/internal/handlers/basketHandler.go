package handlers

import (
	"net/http"

	basket "github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/core/basket"
	"github.com/SandeepMultani/gocommerce/src/backend/basket-srv/pkg/constants"
	httprequest "github.com/SandeepMultani/gocommerce/src/backend/basket-srv/pkg/httpRequest"
	httpresponse "github.com/SandeepMultani/gocommerce/src/backend/basket-srv/pkg/httpResponse"
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

var _ BasketHandler = &basketHandler{}

func NewBasketHandler(basketSrv basket.BasketService) BasketHandler {
	return &basketHandler{
		basketSrv: basketSrv,
	}
}

func (hdl *basketHandler) Get(c *gin.Context) {
	requestId := getRequestHeader(c, constants.HEADER_REQUEST_ID)
	basketId := c.Param("basketId")
	bas, err := hdl.basketSrv.Get(basketId)
	if err != nil {
		c.JSON(http.StatusNotFound, httpresponse.NewHttpErrorResponse(requestId, err.Error()))
		return
	}
	c.JSON(200, httpresponse.NewHttpSuccessResponse(requestId, bas))
}

func (hdl *basketHandler) Create(c *gin.Context) {
	requestId := getRequestHeader(c, constants.HEADER_REQUEST_ID)
	createBasketReq := httprequest.CreateBasketRequest{}
	c.Bind(&createBasketReq)
	bas, err := hdl.basketSrv.Create(createBasketReq.BasketId)
	if err != nil {
		c.JSON(http.StatusNotFound, httpresponse.NewHttpErrorResponse(requestId, err.Error()))
		return
	}
	c.JSON(200, httpresponse.NewHttpSuccessResponseWithMessage(requestId, bas, constants.BASKET_CREATED))
}

func (hdl *basketHandler) Delete(c *gin.Context) {
	requestId := getRequestHeader(c, constants.HEADER_REQUEST_ID)
	basketId := c.Param("basketId")
	err := hdl.basketSrv.Delete(basketId)
	if err != nil {
		c.JSON(http.StatusNotFound, httpresponse.NewHttpErrorResponse(requestId, err.Error()))
		return
	}
	c.JSON(200, httpresponse.NewHttpSuccessResponseWithMessage(requestId, nil, constants.BASKET_DELETED))
}

func (hdl *basketHandler) AddItem(c *gin.Context) {
	requestId := getRequestHeader(c, constants.HEADER_REQUEST_ID)
	basketId := c.Param("basketId")
	p := &basket.Product{}
	c.Bind(p)
	bas, err := hdl.basketSrv.AddItem(basketId, p)
	if err != nil {
		c.JSON(http.StatusNotFound, httpresponse.NewHttpErrorResponse(requestId, err.Error()))
		return
	}
	c.JSON(200, httpresponse.NewHttpSuccessResponseWithMessage(requestId, bas, constants.ITEM_ADDED))
}

func (hdl *basketHandler) RemoveItem(c *gin.Context) {
	requestId := getRequestHeader(c, constants.HEADER_REQUEST_ID)
	basketId := c.Param("basketId")
	removeItemReq := httprequest.RemoveItemRequest{}
	c.Bind(&removeItemReq)
	bas, err := hdl.basketSrv.RemoveItem(basketId, removeItemReq.ProductID)
	if err != nil {
		c.JSON(http.StatusNotFound, httpresponse.NewHttpErrorResponse(requestId, err.Error()))
		return
	}
	c.JSON(200, httpresponse.NewHttpSuccessResponseWithMessage(requestId, bas, constants.ITEM_REMOVED))
}

func getRequestHeader(c *gin.Context, key string) string {
	if values := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}
