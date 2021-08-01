package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	ProductID   string             `json:"product_id" bson:"product_id"`
	Sku         string             `json:"sku" bson:"sku"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float32            `json:"price" bson:"price"`
	UpsertedAt  int64              `json:"upserted_at" bson:"upserted_at"`
}

type ProductRepository interface {
	Get() ([]*Product, error)
	GetById(string) (*Product, error)
	GetBySku(string) (*Product, error)
	Create(*Product) error
}

type ProductService interface {
	Get() ([]*Product, error)
	GetById(string) (*Product, error)
	GetBySku(string) (*Product, error)
	Create(*Product) (*Product, error)
}
