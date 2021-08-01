package handlers

import (
	"net/http"

	product "github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/internal/core/product"
	constants "github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/pkg/constants"
	httpResponse "github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/pkg/httpResponse"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	Get(*gin.Context)
	GetById(*gin.Context)
	GetBySku(*gin.Context)
	Create(*gin.Context)
}

type productHandler struct {
	productSrv product.ProductService
}

var _ ProductHandler = &productHandler{}

func NewProductHandler(productSrv product.ProductService) ProductHandler {
	return &productHandler{
		productSrv: productSrv,
	}
}

func (hdl *productHandler) Get(c *gin.Context) {
	requestId := getRequestHeader(c, constants.HEADER_REQUEST_ID)

	products, err := hdl.productSrv.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpResponse.NewHttpErrorResponse(requestId, err.Error()))
		return
	}

	c.JSON(200, httpResponse.NewHttpSuccessResponse(requestId, products))
}

func (hdl *productHandler) GetById(c *gin.Context) {
	requestId := getRequestHeader(c, constants.HEADER_REQUEST_ID)
	id := c.Param("id")

	product, err := hdl.productSrv.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, httpResponse.NewHttpErrorResponse(requestId, err.Error()))
		return
	}

	c.JSON(200, httpResponse.NewHttpSuccessResponse(requestId, product))
}

func (hdl *productHandler) GetBySku(c *gin.Context) {
	requestId := getRequestHeader(c, constants.HEADER_REQUEST_ID)
	sku := c.Query("sku")

	product, err := hdl.productSrv.GetBySku(sku)
	if err != nil {
		c.JSON(http.StatusNotFound, httpResponse.NewHttpErrorResponse(requestId, err.Error()))
		return
	}

	c.JSON(200, httpResponse.NewHttpSuccessResponse(requestId, product))
}

func (hdl *productHandler) Create(c *gin.Context) {
	requestId := getRequestHeader(c, constants.HEADER_REQUEST_ID)
	product := &product.Product{}
	c.Bind(product)

	newProduct, err := hdl.productSrv.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpResponse.NewHttpErrorResponse(requestId, err.Error()))
		return
	}

	c.JSON(200, httpResponse.NewHttpSuccessResponseWithMessage(requestId, newProduct, constants.PRODUCT_CREATED))
}

func getRequestHeader(c *gin.Context, key string) string {
	if values := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}
