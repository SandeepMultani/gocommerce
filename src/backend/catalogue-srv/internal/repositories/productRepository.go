package repositories

import (
	"context"
	"errors"

	product "github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/internal/core/product"
	"github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/pkg/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepository struct {
	catalogueCollection *mongo.Collection
}

var _ product.ProductRepository = &productRepository{}

func NewProductRepository(catalogueCollection *mongo.Collection) product.ProductRepository {
	return &productRepository{
		catalogueCollection: catalogueCollection,
	}
}

func (repo *productRepository) Get() ([]*product.Product, error) {
	var prods []*product.Product
	cursor, err := repo.catalogueCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return prods, errors.New(constants.DATABASE_OPERATION_ERROR)
	}

	for cursor.Next(context.TODO()) {
		var p product.Product
		cursor.Decode(&p)
		prods = append(prods, &p)
	}

	return prods, nil
}

func (repo *productRepository) GetById(id string) (*product.Product, error) {
	p, exists, err := repo.exists(id)
	if !exists {
		if err == mongo.ErrNoDocuments {
			return &product.Product{}, errors.New(constants.PRODUCT_NOT_FOUND)
		}
		return &product.Product{}, errors.New(constants.DATABASE_OPERATION_ERROR)
	}
	return p, nil
}

func (repo *productRepository) GetBySku(sku string) (*product.Product, error) {
	p, exists, err := repo.skuExists(sku)
	if !exists {
		if err == mongo.ErrNoDocuments {
			return &product.Product{}, errors.New(constants.PRODUCT_NOT_FOUND)
		}
		return &product.Product{}, errors.New(constants.DATABASE_OPERATION_ERROR)
	}
	return p, nil
}

func (repo *productRepository) Create(p *product.Product) error {
	_, exists, _ := repo.skuExists(p.Sku)
	if exists {
		return errors.New(constants.PRODUCT_ALREADY_EXISTS)
	}

	_, err := repo.catalogueCollection.InsertOne(context.TODO(), *p)
	if err != nil {
		return errors.New(constants.DATABASE_OPERATION_ERROR)
	}
	return nil
}

func (repo *productRepository) exists(productId string) (*product.Product, bool, error) {
	p := product.Product{}
	err := repo.catalogueCollection.FindOne(context.TODO(), bson.M{"product_id": productId}).Decode(&p)
	if err == nil {
		return &p, true, nil
	}
	return &product.Product{}, false, err
}

func (repo *productRepository) skuExists(sku string) (*product.Product, bool, error) {
	p := product.Product{}
	err := repo.catalogueCollection.FindOne(context.TODO(), bson.M{"sku": sku}).Decode(&p)
	if err == nil {
		return &p, true, nil
	}
	return &product.Product{}, false, err
}
