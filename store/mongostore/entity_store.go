package mongostore

import (
	"context"
	"errors"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/store/storeutils"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type entityStore[T any] struct {
	collection *mongo.Collection
}

func (s *entityStore[T]) GetAll(opts ...storeutils.QueryOptions) ([]*T, error) {
	entities := []*T{}
	opt := storeutils.QueryOptions{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	err := opt.ApplyTagFieldNames(*new(T), "bson")
	if err != nil {
		return nil, err
	}

	r, err := s.collection.Find(
		context.TODO(),
		opt.Conditions,
		&options.FindOptions{
			Projection: opt.Projection,
		})
	if err != nil {
		return nil, err
	}

	err = r.All(context.TODO(), &entities)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (s *entityStore[T]) GetByID(id string, projection ...storeutils.P) (*T, error) {
	t := new(T)
	proj := storeutils.P{}
	if len(projection) > 0 {
		proj = projection[0]
	}

	err := s.collection.FindOne(context.TODO(), bson.M{"_id": id}, &options.FindOneOptions{
		Projection: proj,
	}).Decode(t)
	if err != nil {
		return nil, errors.New(models.ResourceDoesNotExist)
	}

	return t, nil
}

func (s *entityStore[T]) GetOne(opts ...storeutils.QueryOptions) (*T, error) {
	t := new(T)
	opt := storeutils.QueryOptions{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	err := opt.ApplyTagFieldNames(*new(T), "bson")
	if err != nil {
		return nil, err
	}

	err = s.collection.FindOne(
		context.TODO(),
		opt.Conditions,
		&options.FindOneOptions{
			Projection: opt.Projection,
		}).Decode(t)
	if err != nil {
		return nil, errors.New(models.ResourceDoesNotExist)
	}

	return t, nil
}

func (s *entityStore[T]) Create(entity *T) (*T, error) {
	_, err := s.collection.InsertOne(context.TODO(), entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s *entityStore[T]) Update(entity *T) (*T, error) {
	entityAsMap, err := utils.ToBSON(entity)
	if err != nil {
		return nil, err
	}

	r, err := s.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": entityAsMap["_id"]},
		bson.M{"$set": entity},
	)
	if err != nil {
		return nil, err
	}

	if r.ModifiedCount != 1 {
		return nil, errors.New(models.ResourceUnmodified)
	}

	return entity, nil
}

func (s *entityStore[T]) DeleteAll() error {
	return errors.New("Not implemented")
}

func (s *entityStore[T]) DeleteByID(id string) error {
	r, err := s.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	if r.DeletedCount != 1 {
		return errors.New(models.ResourceUnmodified)
	}

	return nil
}
