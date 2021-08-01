package repositories

import (
	"context"
	"sync"
	"time"

	"github.com/SandeepMultani/gocommerce/src/backend/catalogue-srv/pkg/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
)

func getMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		// Database Config
		clientOptions := options.Client().ApplyURI(constants.MONGO_CONNECTION_STRING)
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context.Background(), readpref.Primary())
		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func GetCatalogueCollection() (*mongo.Collection, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(constants.MONGO_CATALOGUE_DB).Collection(constants.MONGO_CATALOGUE_COLLECTION)
	return collection, nil
}
