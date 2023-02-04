package api

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/logger"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/store/storeutils"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"
)

func (api *Api) registerUserRoutes() {
	ur := api.routes.user
	ur.Handle("/register", api.handleAPI(register)).Methods("POST", "OPTIONS")
	ur.Handle("/login", api.handleAPI(login)).Methods("POST", "OPTIONS")
	ur.Handle("/oauth/{provider}/login", api.handleAPI(oAuthLogin)).Methods("GET", "OPTIONS")
	ur.Handle("/oauth/{provider}/callback", api.handleAPI(oAuthCallback)).Methods("GET", "OPTIONS")
	ur.Handle("/send-verification-email", api.handleAPI(sendVerificationEmail)).Methods("POST", "OPTIONS")
	ur.Handle("/verify-email", api.handleAPI(verifyEmail)).Methods("PATCH", "OPTIONS")
	ur.Handle("/change-password", api.handleAuthRequired(changePassword)).Methods("PATCH", "OPTIONS")
	ur.Handle("/send-reset-password-link", api.handleAuthRequired(sendResetPasswordLink)).Methods("POST", "OPTIONS")
	ur.Handle("/reset-password", api.handleAuthRequired(resetPassword)).Methods("PATCH", "OPTIONS")
	ur.Handle("/logout", api.handleAuthRequired(logout)).Methods("POST", "OPTIONS")
	ur.Handle("/teams", api.handleAuthRequired(getCurrentUserTeams)).Methods("GET", "OPTIONS")
	ur.Handle("/", api.handleAuthRequired(getCurrentUser)).Methods("GET", "OPTIONS")
}

func register(c *apiContext, w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	isValid, registerData := models.ValidateRegisterData(reqBody)
	if !isValid {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	savedUser, _ := c.app.UserStore.GetOne(storeutils.QueryOptions{
		Conditions: storeutils.M{"Username": registerData.Username},
	})
	if savedUser != nil {
		sendErrorResponse("Username already exists", http.StatusBadRequest, w)
		return
	}

	savedUser, _ = c.app.UserStore.GetOne(storeutils.QueryOptions{
		Conditions: storeutils.M{"Email": registerData.Email},
	})
	if savedUser != nil {
		sendErrorResponse("Email already exists", http.StatusBadRequest, w)
		return
	}

	user := new(models.User)
	err := user.Presave(registerData)
	if err != nil {
		logger.Error("register user.Presave", err)
		sendErrorResponse(models.SomethingWentWrong, http.StatusInternalServerError, w)
		return
	}

	u, err := c.app.UserStore.Create(user)
	if err != nil {
		logger.Error("register UserStore.Create", err)
		sendErrorResponse(models.SomethingWentWrong, http.StatusInternalServerError, w)
		return
	}

	t, err := c.app.TokenStore.Create(&models.Token{
		ID:        utils.NewLongID(),
		Type:      models.SessionToken,
		ExpiresAt: time.Now().Add(models.SessionTokenExpiryDuration),
		OwnerID:   u.ID,
	})
	if err != nil {
		logger.Error("register TokenStore.Create", err)
		sendErrorResponse(models.SomethingWentWrong, http.StatusInternalServerError, w)
		return
	}

	setCookies(w, u.ID, t.ID)
	sendJSONResponse(u, http.StatusOK, w)
}

func login(c *apiContext, w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	isValid, loginData := models.ValidateLoginData(reqBody)
	if !isValid {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	var savedUser *models.User
	var err error
	var msg string

	if loginData.Username != "" {
		savedUser, err = c.app.UserStore.GetOne(storeutils.QueryOptions{
			Conditions: storeutils.M{"Username": loginData.Username},
		})
		msg = "Invalid Username"
	} else {
		savedUser, err = c.app.UserStore.GetOne(storeutils.QueryOptions{
			Conditions: storeutils.M{"Email": loginData.Email},
		})
		msg = "Invalid Email"
	}

	if err != nil {
		sendErrorResponse(msg, http.StatusBadRequest, w)
		return
	}

	if !models.IsCorrectPassword(savedUser.Password, loginData.Password) {
		sendErrorResponse("Invalid password", http.StatusBadRequest, w)
		return
	}

	t, err := c.app.TokenStore.Create(&models.Token{
		ID:        utils.NewLongID(),
		Type:      models.SessionToken,
		ExpiresAt: time.Now().Add(models.SessionTokenExpiryDuration),
		OwnerID:   savedUser.ID,
	})
	if err != nil {
		logger.Error("login TokenStore.Create", err)
		sendErrorResponse(models.SomethingWentWrong, http.StatusInternalServerError, w)
		return
	}

	setCookies(w, savedUser.ID, t.ID)
	sendJSONResponse(savedUser, http.StatusOK, w)
}

func oAuthLogin(c *apiContext, w http.ResponseWriter, req *http.Request) {
	p := mux.Vars(req)["provider"]
	provider, ok := oAuthProviders[p]
	if !ok {
		sendErrorResponse("Unsupported OAuth Provider", http.StatusBadRequest, w)
		return
	}

	oAuthState := utils.NewID()
	stateCookie := http.Cookie{
		Name:    "oauthstate",
		Value:   oAuthState,
		Expires: time.Now().Add(time.Hour),
	}
	refererCookie := http.Cookie{
		Name:    "referer",
		Value:   req.Referer(),
		Expires: time.Now().Add(time.Hour),
	}

	http.SetCookie(w, &stateCookie)
	http.SetCookie(w, &refererCookie)
	u := provider.AuthCodeURL(oAuthState)

	http.Redirect(w, req, u, http.StatusTemporaryRedirect)
}

func oAuthCallback(c *apiContext, w http.ResponseWriter, req *http.Request) {
	oAuthState, err := req.Cookie("oauthstate")
	if oAuthState == nil || err != nil {
		http.Redirect(w, req, "/", http.StatusBadRequest)
		return
	}

	if oAuthState.Value != req.FormValue("state") {
		http.Redirect(w, req, "/", http.StatusBadRequest)
		return
	}

	p := mux.Vars(req)["provider"]
	provider, ok := oAuthProviders[p]
	if !ok {
		http.Redirect(w, req, "/", http.StatusBadRequest)
		return
	}

	user, err := provider.getUser(req.FormValue("code"))
	if err != nil {
		logger.Error("OAuthProvider.getUser", err)
		http.Redirect(w, req, "/", http.StatusInternalServerError)
		return
	}

	savedUser, err := c.app.UserStore.GetByID(user.ID)
	if savedUser == nil && err != nil {
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		createdUser, err := c.app.UserStore.Create(user)
		if err != nil {
			logger.Error("Save oauth user", err)
			http.Redirect(w, req, "/", http.StatusBadRequest)
			return
		}
		user = createdUser
	} else {
		user = savedUser
	}

	t, err := c.app.TokenStore.Create(&models.Token{
		ID:        utils.NewLongID(),
		Type:      models.SessionToken,
		ExpiresAt: time.Now().Add(models.SessionTokenExpiryDuration),
		OwnerID:   user.ID,
	})
	if err != nil {
		logger.Error("login TokenStore.Create", err)
		sendErrorResponse(models.SomethingWentWrong, http.StatusInternalServerError, w)
		return
	}

	setCookies(w, user.ID, t.ID)

	referer, _ := req.Cookie("referer")
	http.Redirect(w, req, referer.Value, http.StatusPermanentRedirect)
}

func sendVerificationEmail(c *apiContext, w http.ResponseWriter, req *http.Request) {
}

func verifyEmail(c *apiContext, w http.ResponseWriter, req *http.Request) {
}

func changePassword(c *apiContext, w http.ResponseWriter, req *http.Request) {
}

func sendResetPasswordLink(c *apiContext, w http.ResponseWriter, req *http.Request) {
}

func resetPassword(c *apiContext, w http.ResponseWriter, req *http.Request) {
}

func logout(c *apiContext, w http.ResponseWriter, req *http.Request) {
	var err error

	if c.token.Type == models.SessionToken && c.token.OwnerID == c.user.ID {
		err = c.app.TokenStore.DeleteByID(c.token.ID)
		if err != nil {
			logger.Error("logout TokenStore.DeleteByID", err)
			sendErrorResponse(models.SomethingWentWrong, http.StatusInternalServerError, w)
			return
		}
	}

	removeCookies(w)
	w.WriteHeader(http.StatusNoContent)
}

func getCurrentUser(c *apiContext, w http.ResponseWriter, req *http.Request) {
	sendJSONResponse(c.user, http.StatusOK, w)
}

func getCurrentUserTeams(c *apiContext, w http.ResponseWriter, req *http.Request) {
	teams, err := c.app.TeamStore.GetByUserID(c.user.ID, true)
	if err != nil {
		logger.Error("userAPI.getCurrentUserTeams", err)
		sendErrorResponse("Could not get teams", http.StatusBadRequest, w)
		return
	}

	sendJSONResponse(teams, http.StatusOK, w)
}
