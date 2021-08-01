package basket

import "time"

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
	newBasket := &Basket{
		ID:         basketId,
		Items:      map[string]BasketItem{},
		Total:      0,
		UpsertedAt: time.Now().Unix(),
	}
	err := basketSrv.repository.Create(newBasket)
	if err != nil {
		return &Basket{}, err
	}
	return newBasket, nil
}

func (basketSrv *basketService) Delete(basketId string) error {
	return basketSrv.repository.Delete(basketId)
}

func (basketSrv *basketService) AddItem(basketId string, product *Product) (*Basket, error) {
	b, err := basketSrv.repository.Get(basketId)
	if err != nil {
		return &Basket{}, err
	}

	if basketItem, ok := b.Items[product.ProductID]; ok {
		basketItem.Quantity++
		basketItem.Total = basketItem.Price * float32(basketItem.Quantity)
		b.Items[product.ProductID] = basketItem
	} else {
		newBasketItem := &BasketItem{
			ProductID:   product.ProductID,
			Sku:         product.Sku,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    1,
			Total:       product.Price,
		}
		b.Items[product.ProductID] = *newBasketItem
	}

	b.calculateTotal()

	err = basketSrv.repository.Update(b)
	if err != nil {
		return &Basket{}, err
	}
	return b, nil
}

func (basketSrv *basketService) RemoveItem(basketId, productId string) (*Basket, error) {
	b, err := basketSrv.repository.Get(basketId)
	if err != nil {
		return &Basket{}, err
	}

	if basketItem, ok := b.Items[productId]; ok {
		basketItem.Quantity--
		if basketItem.Quantity == 0 {
			delete(b.Items, productId)
		} else {
			basketItem.Total = basketItem.Price * float32(basketItem.Quantity)
			b.Items[productId] = basketItem
		}
	}

	b.calculateTotal()

	err = basketSrv.repository.Update(b)
	if err != nil {
		return &Basket{}, err
	}
	return b, nil
}

func (b *Basket) calculateTotal() {
	var total float32
	for k := range b.Items {
		total = total + b.Items[k].Total
	}
	b.Total = total
}
