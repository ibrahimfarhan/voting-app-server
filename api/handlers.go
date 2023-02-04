package api

import (
	"net/http"
	"time"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/config"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/app"
)

type apiHandler struct {
	app          *app.App
	handleFunc   func(*apiContext, http.ResponseWriter, *http.Request)
	authRequired bool
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", config.Env.AllowedOrigins)
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (h *apiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w)

	if req.Method == "OPTIONS" {
		return
	}

	context := &apiContext{
		app:       h.app,
		path:      req.URL.Path,
		user:      nil,
		token:     nil,
		ipAddress: utils.GetIPAddress(req),
		method:    req.Method,
	}

	if !h.authRequired {
		h.handleFunc(context, w, req)
		return
	}

	sessionCookie, _ := req.Cookie("session")
	if sessionCookie == nil {
		sendErrorResponse(models.AuthRequired, http.StatusUnauthorized, w)
		return
	}

	tokenID := sessionCookie.Value

	t, err := h.app.TokenStore.GetByID(tokenID)
	if err != nil {
		sendErrorResponse(models.AuthRequired, http.StatusUnauthorized, w)
		return
	}

	if time.Now().After(t.ExpiresAt) {
		h.app.TokenStore.DeleteByID(t.ID)
		sendErrorResponse(models.AuthRequired, http.StatusUnauthorized, w)
		return
	}

	user, err := h.app.UserStore.GetByID(t.OwnerID)
	if err != nil {
		sendErrorResponse(models.AuthRequired, http.StatusUnauthorized, w)
		return
	}

	context.token = t
	context.user = user

	h.handleFunc(context, w, req)
}
