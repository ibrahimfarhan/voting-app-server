package mongostore

import (
	"context"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const maxConnectionRetryTimes = 5

type MongoConnector struct {
	client *mongo.Client
	DB     *mongo.Database
}

func NewMongoConnector() *MongoConnector {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Env.DBConnectionString))
	if err != nil {
		panic(err)
	}

	connector := &MongoConnector{
		client: client,
		DB:     client.Database(config.Env.DBName),
	}

	return connector
}

func (c *MongoConnector) Open(ctx context.Context) error {
	var err error
	for i := 0; i < maxConnectionRetryTimes; i++ {
		err = c.client.Connect(ctx)
		if err == nil {
			return nil
		}
	}

	return err
}

func (c *MongoConnector) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}
