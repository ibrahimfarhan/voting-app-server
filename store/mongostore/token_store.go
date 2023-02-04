package mongostore

import (
	"context"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenStore struct {
	entityStore[models.Token]
}

func NewTokenStore(collection *mongo.Collection) *TokenStore {
	return &TokenStore{
		entityStore: entityStore[models.Token]{
			collection: collection,
		},
	}
}

func (s *TokenStore) RegisterIndexes() error {
	var expiry int32 = 0
	background := true
	name := "tokenExpiryTTL"
	_, err := s.collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys: bson.M{"expiresAt": 1},
			Options: &options.IndexOptions{
				ExpireAfterSeconds: &expiry,
				Background:         &background,
				Name:               &name,
			},
		},
		{
			Keys: bson.M{"ownerId": 1},
		},
	})
	return err
}
