package api

import (
	"net/http"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/app"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/config"

	"github.com/gorilla/mux"
)

type Routes struct {
	root   *mux.Router
	user   *mux.Router
	team   *mux.Router
	voting *mux.Router
}

type Api struct {
	app    *app.App
	routes *Routes
}

func InitApi(a *app.App) {
	api := &Api{routes: &Routes{}}
	api.app = a
	api.routes.root = a.RootRouter.PathPrefix("/api/v1").Subrouter()
	api.routes.user = api.routes.root.PathPrefix("/user").Subrouter()
	api.routes.team = api.routes.root.PathPrefix("/team").Subrouter()
	api.routes.voting = api.routes.root.PathPrefix("/voting").Subrouter()

	api.registerUserRoutes()
	api.registerTeamRoutes()
	api.registerVotingRoutes()

	if config.Env.ServeStaticFiles == "true" {
		a.RegisterStaticRoutes()
	} else {
		a.RootRouter.HandleFunc("/", (func(w http.ResponseWriter, r *http.Request) {
			sendJSONResponse(map[string]bool{"started": true}, http.StatusOK, w)
		}))

		a.RootRouter.PathPrefix("/{any:.*}").HandlerFunc(api.handle404)
	}

	api.routes.root.PathPrefix("/{any:.*}").HandlerFunc(api.handle404)
}

func (api *Api) handleAPI(fn func(*apiContext, http.ResponseWriter, *http.Request)) http.Handler {
	return &apiHandler{
		app:          api.app,
		handleFunc:   fn,
		authRequired: false,
	}
}

func (api *Api) handleAuthRequired(fn func(*apiContext, http.ResponseWriter, *http.Request)) http.Handler {
	return &apiHandler{
		app:          api.app,
		handleFunc:   fn,
		authRequired: true,
	}
}

func (api *Api) handle404(w http.ResponseWriter, req *http.Request) {
	sendErrorResponse("Not Found", http.StatusNotFound, w)
}
