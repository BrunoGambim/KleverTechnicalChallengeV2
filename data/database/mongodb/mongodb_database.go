package mongodb_database_factory

import (
	"context"
	"os"
	mongodb_errors "relembrando_2/data/database/mongodb/errors"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseInstance *MongodbDatabase

var databaseInstanceError error

var mongoOnce sync.Once

type MongodbDatabase struct {
	mongoDatabase *mongo.Database
}

func NewMongoDatabase(ctx context.Context) (*MongodbDatabase, error) {
	mongoOnce.Do(func() {
		databaseName := os.Getenv("MONGODB_DATABASE_NAME")

		client, err := getMongoClient(ctx)

		databaseInstanceError = err
		if err == nil {
			database := client.Database(databaseName)
			databaseInstance = &MongodbDatabase{
				mongoDatabase: database,
			}
		}
	})
	return databaseInstance, mongodb_errors.Handle(databaseInstanceError)
}

func (database MongodbDatabase) GetCollection(collection string) *mongo.Collection {
	return database.mongoDatabase.Collection(collection)
}

func getMongoClient(ctx context.Context) (*mongo.Client, error) {
	databaseConnection := os.Getenv("MONGODB_CONNECTION")
	clientOptions := options.Client().ApplyURI(databaseConnection)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	return client, err
}
