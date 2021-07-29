package basket

type basketService struct {
	repository BasketRepository
}

var _ BasketService = &basketService{}

func NewBasketService(repository BasketRepository) BasketService {
	return &basketService{
		repository: repository,
	}
}

func (basketSrv *basketService) Get(basketId string) (*Basket, error) {
	return basketSrv.repository.Get(basketId)
}

func (basketSrv *basketService) Create(basketId string) (*Basket, error) {
	return basketSrv.repository.Create(basketId)
}

func (basketSrv *basketService) Delete(basketId string) error {
	return basketSrv.repository.Delete(basketId)
}

func (basketSrv *basketService) AddItem(basketId, productId string) (*Basket, error) {
	return basketSrv.repository.AddItem(basketId, productId)
}

func (basketSrv *basketService) RemoveItem(basketId, productId string) (*Basket, error) {
	return basketSrv.repository.RemoveItem(basketId, productId)
}
