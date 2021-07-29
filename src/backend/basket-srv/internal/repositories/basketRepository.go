package repositories

import (
	"errors"

	basket "github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/core/basket"
)

var (
	_       basket.BasketRepository = &basketRepository{}
	baskets []basket.Basket         = []basket.Basket{
		{
			ID: "123",
			Items: []basket.BasketItem{
				{
					ID:       "234",
					Name:     "Apple",
					Price:    2.50,
					Quantity: 1,
				},
			},
		},
	}
)

type basketRepository struct {
}

func NewBasketRepository() basket.BasketRepository {
	return &basketRepository{}
}

func (repo *basketRepository) Get(basketId string) (*basket.Basket, error) {
	bas, exists := repo.exists(basketId)
	if !exists {
		return &basket.Basket{}, errors.New("not found")
	}
	return bas, nil
}

func (repo *basketRepository) Create(basketId string) (*basket.Basket, error) {
	bas, exists := repo.exists(basketId)
	if exists {
		return bas, errors.New("already exists")
	}

	newBasket := basket.Basket{
		ID:    basketId,
		Items: []basket.BasketItem{},
	}
	baskets = append(baskets, newBasket)

	return &newBasket, nil
}

func (repo *basketRepository) Delete(basketId string) error {
	for i, v := range baskets {
		if v.ID == basketId {
			baskets = append(baskets[:i], baskets[i+1:]...)
			break
		}
	}
	return nil
}

func (repo *basketRepository) AddItem(basketId string, productId string) (*basket.Basket, error) {
	bas, exists := repo.exists(basketId)
	if !exists {
		return &basket.Basket{}, errors.New("not found")
	}

	bas.Items = append(bas.Items, basket.BasketItem{
		ID:       productId,
		Quantity: 1,
	})

	repo.update(basketId, *bas)

	return bas, nil
}

func (repo *basketRepository) RemoveItem(basketId string, productId string) (*basket.Basket, error) {
	bas, exists := repo.exists(basketId)
	if !exists {
		return &basket.Basket{}, errors.New("not found")
	}

	for i, v := range bas.Items {
		if v.ID == productId {
			bas.Items = append(bas.Items[:i], bas.Items[i+1:]...)
			break
		}
	}

	repo.update(basketId, *bas)

	return bas, nil
}

func (repo *basketRepository) exists(basketId string) (*basket.Basket, bool) {
	for _, b := range baskets {
		if b.ID == basketId {
			return &b, true
		}
	}
	return &basket.Basket{}, false
}

func (repo *basketRepository) update(basketId string, updatedBasket basket.Basket) error {
	for i, b := range baskets {
		if b.ID == basketId {
			baskets[i] = updatedBasket
			break
		}
	}
	return nil
}
