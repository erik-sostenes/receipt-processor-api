package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClientFactory struct{}

func (*MongoClientFactory) createConnectClient(uri, databaseName string) (mongoDB *mongo.Database, err error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return
	}

	mongoDB = client.Database(databaseName)

	var result bson.M
	if err := mongoDB.RunCommand(context.Background(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, err
	}

	return
}

func (m *MongoClientFactory) CreateClient(uri, databaseName string) (*mongo.Database, error) {
	return m.createConnectClient(uri, databaseName)
}
