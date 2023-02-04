package app

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/config"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"
)

var serverDirPath = utils.GetExecutableDirPath()
var indexPath = filepath.Join(serverDirPath, "../client")

func (a *App) RegisterStaticRoutes() {
	staticHandler := staticFilesHandler(http.FileServer(http.Dir(indexPath)))

	a.RootRouter.PathPrefix("/static/").Handler(staticHandler)

	a.RootRouter.PathPrefix("/{any:.*}").HandlerFunc(indexHandler).Methods("GET", "OPTIONS")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", config.Env.AllowedOrigins)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if req.Method == "OPTIONS" {
		return
	}

	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Expires", "0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	http.ServeFile(w, req, filepath.Join(indexPath, "index.html"))
}

// A wrapper to allow for setting content-type of static files.
func staticFilesHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		var contentType string

		if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else {
			contentType = "text/plain"
		}

		w.Header().Set("Content-Type", contentType)
		handler.ServeHTTP(w, r)
	})
}
