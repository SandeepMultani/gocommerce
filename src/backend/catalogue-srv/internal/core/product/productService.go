package product

import (
	"time"

	uuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productService struct {
	repository ProductRepository
}

var _ ProductService = &productService{}

func NewProductService(repository ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (srv *productService) Get() ([]*Product, error) {
	return srv.repository.Get()
}

func (srv *productService) GetById(id string) (*Product, error) {
	return srv.repository.GetById(id)
}

func (srv *productService) GetBySku(sku string) (*Product, error) {
	return srv.repository.GetBySku(sku)
}

func (srv *productService) Create(p *Product) (*Product, error) {
	p.ID = primitive.NewObjectID()
	p.ProductID = uuid.NewString()
	p.UpsertedAt = time.Now().Unix()
	if err := srv.repository.Create(p); err != nil {
		return &Product{}, err
	}
	return p, nil
}
