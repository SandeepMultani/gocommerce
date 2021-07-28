package repositories

import (
	"errors"

	product "github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/internal/core/product"
)

type productRepository struct{}

var (
	_        product.ProductRepository = &productRepository{}
	products []product.Product         = []product.Product{
		{
			ID:   "123",
			Name: "Product 123",
			Sku:  "sku123",
		},
		{
			ID:   "456",
			Name: "Product 456",
			Sku:  "sku456",
		},
		{
			ID:   "789",
			Name: "Product 789",
			Sku:  "sku789",
		},
	}
)

func NewProductRepository() product.ProductRepository {
	return &productRepository{}
}

func (srv *productRepository) Get(id string) (*product.Product, error) {
	for _, v := range products {
		if v.ID == id {
			return &v, nil
		}
	}
	return &product.Product{}, errors.New("not found")
}

func (srv *productRepository) GetBySku(sku string) (*product.Product, error) {
	for _, v := range products {
		if v.Sku == sku {
			return &v, nil
		}
	}
	return &product.Product{}, errors.New("not found")
}
