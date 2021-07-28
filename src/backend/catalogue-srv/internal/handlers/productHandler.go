package handlers

import (
	product "github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/internal/core/product"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	Get(*gin.Context)
	GetBySku(*gin.Context)
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
	id := c.Param("id")
	product, err := hdl.productSrv.Get(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, product)
}

func (hdl *productHandler) GetBySku(c *gin.Context) {
	sku := c.Query("sku")
	product, err := hdl.productSrv.GetBySku(sku)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, product)
}
