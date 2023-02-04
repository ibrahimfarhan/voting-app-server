package api

import (
	"github.com/ibrahimfarhan/voting-app/voting-app-server/app"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
)

// Context is passed to the api handlers to give access to the stores and user session.
type apiContext struct {
	app       *app.App
	user      *models.User
	token     *models.Token
	ipAddress string
	path      string
	method    string
}
