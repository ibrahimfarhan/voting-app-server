package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/config"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type oAuthProvider struct {
	oauth2.Config
	userApiURL               string
	mapProviderUserToAppUser func(data map[string]interface{}) *models.User
}

var googleOAuth = &oAuthProvider{
	Config: oauth2.Config{
		RedirectURL:  strings.Replace(config.Env.OAuthRedirectURL, "provider", "google", 1),
		ClientID:     config.Env.GoogleOAuthClientID,
		ClientSecret: config.Env.GoogleOAuthClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	},
	userApiURL: "https://www.googleapis.com/oauth2/v2/userinfo?access_token=",
	mapProviderUserToAppUser: func(data map[string]interface{}) *models.User {
		return &models.User{
			ID:         utils.NewID(),
			OAuthID:    data["id"].(string),
			Name:       data["name"].(string),
			PictureURL: data["picture"].(string),
			Email:      data["email"].(string),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
	},
}

var oAuthProviders = map[string]*oAuthProvider{
	"google": googleOAuth,
}

func (p oAuthProvider) getUser(code string) (*models.User, error) {
	token, err := p.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(p.userApiURL + token.AccessToken)
	if err != nil {
		return nil, err
	}

	data := utils.MapFromBody(res.Body)

	return p.mapProviderUserToAppUser(data), nil
}
