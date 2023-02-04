package interfaces

import "context"

type DBConnector interface {
	Open(ctx context.Context) error
	Close(ctx context.Context) error
}
