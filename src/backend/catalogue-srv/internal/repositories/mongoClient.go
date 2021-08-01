package repositories

import (
	"context"
	"sync"

	"github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/pkg/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
)

func getMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(constants.MONGO_CONNECTION_STRING)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
