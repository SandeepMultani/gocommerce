package repositories

import (
	"sync"

	"github.com/SandeepMultani/gocommerce/src/backend/basket-srv/pkg/constants"
	"github.com/go-redis/redis"
)

var (
	clientInstance      *redis.Client
	clientInstanceError error
	redisOnce           sync.Once
)

func GetRedisClient() (*redis.Client, error) {
	redisOnce.Do(func() {
		options := &redis.Options{
			Addr:     constants.REDIS_CONNECTION_STRING,
			Password: constants.REDIS_PASSWORD,
			DB:       constants.REDIS_DATABASE,
		}
		client := redis.NewClient(options)

		_, err := client.Ping().Result()
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
