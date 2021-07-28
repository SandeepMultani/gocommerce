package product

type productService struct {
	repository ProductRepository
}

var _ ProductService = &productService{}

func NewProductService(repository ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (srv *productService) Get(id string) (*Product, error) {
	return srv.repository.Get(id)
}

func (srv *productService) GetBySku(sku string) (*Product, error) {
	return srv.repository.GetBySku(sku)
}
