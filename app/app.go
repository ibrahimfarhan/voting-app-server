package app

import (
	"context"
	"net/http"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/config"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/realtime"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/store"

	"github.com/gorilla/mux"
)

type App struct {
	server     *http.Server
	RootRouter *mux.Router
	store.Store
	VotingHub *realtime.VotingHub
}

var appCount = 0

func NewApp() *App {
	appCount++
	if appCount > 1 {
		panic("Only one app can be initialized at a time.")
	}

	app := new(App)
	app.RootRouter = mux.NewRouter()
	app.server = NewServer(app.RootRouter, config.Env.Port)
	app.Store = store.InitStore(context.TODO(), config.Env.DBType)
	app.VotingHub = realtime.NewVotingHub()

	go app.VotingHub.Run()

	return app
}
