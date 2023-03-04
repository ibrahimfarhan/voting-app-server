package interfaces

import "github.com/ibrahimfarhan/voting-app/voting-app-server/store/storeutils"

type entityStore[T any] interface {
	RegisterIndexes() error
	GetAll(opts ...storeutils.QueryOptions) ([]*T, error)
	Count(opts ...storeutils.QueryOptions) (int64, error)
	GetByID(id string, projection ...storeutils.P) (*T, error)
	GetOne(opts ...storeutils.QueryOptions) (*T, error)
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	DeleteAll() error
	DeleteByID(id string) error
}
