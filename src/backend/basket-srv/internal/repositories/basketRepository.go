package repositories

import (
	"encoding/json"
	"errors"
	"time"

	basket "github.com/SandeepMultani/gocommerce/src/backend/basket-srv/internal/core/basket"
	"github.com/SandeepMultani/gocommerce/src/backend/basket-srv/pkg/constants"
	"github.com/go-redis/redis"
)

const timeToLive = time.Duration(constants.REDIS_TTL_DAYS*24) * time.Hour

var _ basket.BasketRepository = &basketRepository{}

type basketRepository struct {
	redis *redis.Client
}

func NewBasketRepository(redis *redis.Client) basket.BasketRepository {
	return &basketRepository{
		redis: redis,
	}
}

func (repo *basketRepository) Get(basketId string) (*basket.Basket, error) {
	bas, err := repo.get(basketId)
	if err != nil {
		return &basket.Basket{}, errors.New(constants.BASKET_NOT_FOUND)
	}
	return bas, nil
}

func (repo *basketRepository) Create(newBasket *basket.Basket) error {
	_, err := repo.get(newBasket.ID)
	if err == nil {
		return errors.New(constants.BASKET_ALREADY_EXISTS)
	}

	err = repo.set(newBasket.ID, *newBasket)
	if err != nil {
		return errors.New(constants.DATABASE_OPERATION_ERROR)
	}
	return nil
}

func (repo *basketRepository) Delete(basketId string) error {
	err := repo.redis.Del(basketId).Err()
	if err != nil {
		return errors.New(constants.DATABASE_OPERATION_ERROR)
	}
	return nil
}

func (repo *basketRepository) Update(b *basket.Basket) error {
	err := repo.set(b.ID, *b)
	if err != nil {
		return errors.New(constants.DATABASE_OPERATION_ERROR)
	}
	return nil
}

func (repo *basketRepository) set(key string, b basket.Basket) error {
	json, err := json.Marshal(b)
	if err != nil {
		return err
	}

	err = repo.redis.Set(key, json, timeToLive).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *basketRepository) get(key string) (*basket.Basket, error) {
	val, err := repo.redis.Get(key).Result()
	if err != nil {
		return &basket.Basket{}, err
	}

	b := &basket.Basket{}
	err = json.Unmarshal([]byte(val), &b)
	if err != nil {
		return &basket.Basket{}, err
	}

	return b, nil
}
