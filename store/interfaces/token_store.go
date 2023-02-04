package interfaces

import "github.com/ibrahimfarhan/voting-app/voting-app-server/models"

type TokenStore interface {
	entityStore[models.Token]
}
