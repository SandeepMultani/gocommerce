package repositories

import (
	"context"
	"errors"

	product "github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/internal/core/product"
	"github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/pkg/constants"
)

type productRepository struct{}

var (
	_        product.ProductRepository = &productRepository{}
	products []product.Product         = []product.Product{
		{
			ProductID:   "123",
			Name:        "Product 123",
			Sku:         "sku123",
			Description: "",
			Price:       4.50,
			UpsertedAt:  1627771372,
		},
		{
			ProductID: "456",
			Name:      "Product 456",
			Sku:       "sku456",
		},
		{
			ProductID: "789",
			Name:      "Product 789",
			Sku:       "sku789",
		},
	}
)

func NewProductRepository() product.ProductRepository {
	return &productRepository{}
}

func (repo *productRepository) Get() ([]*product.Product, error) {
	var prods []*product.Product
	for _, p := range products {
		prods = append(prods, &product.Product{
			ProductID:   p.ProductID,
			Sku:         p.Sku,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			UpsertedAt:  p.UpsertedAt,
		})
	}
	return prods, nil
}

func (repo *productRepository) GetById(id string) (*product.Product, error) {
	pro, exists := repo.exists(id)
	if !exists {
		return &product.Product{}, errors.New(constants.PRODUCT_NOT_FOUND)
	}
	return pro, nil
}

func (repo *productRepository) GetBySku(sku string) (*product.Product, error) {
	pro, exists := repo.skuExists(sku)
	if !exists {
		return &product.Product{}, errors.New(constants.PRODUCT_NOT_FOUND)
	}
	return pro, nil
}

func (repo *productRepository) Create(p *product.Product) error {
	_, exists := repo.skuExists(p.Sku)
	if exists {
		return errors.New(constants.PRODUCT_ALREADY_EXISTS)
	}

	client, err := getMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(constants.MONGO_CATALOGUE_DB).Collection(constants.MONGO_CATALOGUE_COLLECTION)
	_, err = collection.InsertOne(context.TODO(), *p)
	if err != nil {
		return err
	}
	return nil
}

func (repo *productRepository) exists(id string) (*product.Product, bool) {
	for _, v := range products {
		if v.ProductID == id {
			return &v, true
		}
	}
	return &product.Product{}, false
}

func (repo *productRepository) skuExists(sku string) (*product.Product, bool) {
	for _, v := range products {
		if v.Sku == sku {
			return &v, true
		}
	}
	return &product.Product{}, false
}
