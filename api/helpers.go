package api

import (
	"encoding/json"
	"net/http"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/config"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
)

func getCookieProps() (http.SameSite, bool) {
	var sameSite http.SameSite
	secure := false

	if config.Env.Env == config.SupportedEnvs.Development {
		sameSite = http.SameSiteLaxMode
	} else {
		sameSite = http.SameSiteNoneMode
		secure = true
	}

	return sameSite, secure
}

func setCookies(w http.ResponseWriter, userId, sessionToken string) {
	sameSite, secure := getCookieProps()

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionToken,
		MaxAge:   3600 * 24 * 14,
		Path:     "/",
		SameSite: sameSite,
		Secure:   secure,
		HttpOnly: true,
	})
}

func removeCookies(w http.ResponseWriter) {
	sameSite, secure := getCookieProps()

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		SameSite: sameSite,
		Secure:   secure,
		HttpOnly: true,
	})
}

func getCookieValues(req *http.Request) (string, string) {
	sessionCookie, err := req.Cookie("")

	userCookie, _err := req.Cookie("")

	if err != nil || _err != nil {
		return "", ""
	}

	return sessionCookie.Value, userCookie.Value
}

func sendJSONResponse(responseData interface{}, status int, w http.ResponseWriter) {
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		status = http.StatusInternalServerError
		jsonData = nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if string(jsonData) != "null" {
		w.Write(jsonData)
	}
}

func sendTextResponse(responseData string, status int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(responseData))
}

func sendErrorResponse(msg string, status int, w http.ResponseWriter) {
	sendJSONResponse(models.NewServerError(msg, status), status, w)
}
