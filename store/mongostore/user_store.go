package mongostore

import (
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore struct {
	entityStore[models.User]
}

func NewUserStore(collection *mongo.Collection) *UserStore {
	return &UserStore{
		entityStore: entityStore[models.User]{
			collection: collection,
		},
	}
}

func (s *UserStore) RegisterIndexes() error {
	return nil
}
