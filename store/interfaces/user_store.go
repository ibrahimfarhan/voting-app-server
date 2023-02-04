package interfaces

import "github.com/ibrahimfarhan/voting-app/voting-app-server/models"

type UserStore interface {
	entityStore[models.User]
}
