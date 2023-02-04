package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"
	"github.com/joho/godotenv"
)

var (
	SupportedEnvs = struct {
		Development string
		Production  string
	}{
		"development",
		"production",
	}

	requiredVars = []string{
		"PORT",
		"ENV",
		"DB_TYPE",
		"DB_CONNECTION_STRING",
		"DB_NAME",
		"SERVE_STATIC_FILES",
		"ALLOWED_ORIGINS",
		"GOOGLE_OAUTH_CLIENT_ID",
		"GOOGLE_OAUTH_CLIENT_SECRET",
		"OAUTH_REDIRECT_URL",
	}

	defaultVars = map[string]string{
		"API_VERSION":        "v1",
		"SERVE_STATIC_FILES": "false",
		"ENV":                "production",
	}

	envVars = getEnvVariables()
	Env     = struct {
		Port                    string
		Env                     string
		DBType                  string
		DBConnectionString      string
		DBName                  string
		AllowedOrigins          string
		ServeStaticFiles        string
		GoogleOAuthClientID     string
		GoogleOAuthClientSecret string
		OAuthRedirectURL        string
	}{
		envVars["PORT"],
		envVars["ENV"],
		envVars["DB_TYPE"],
		envVars["DB_CONNECTION_STRING"],
		envVars["DB_NAME"],
		envVars["ALLOWED_ORIGINS"],
		envVars["SERVE_STATIC_FILES"],
		envVars["GOOGLE_OAUTH_CLIENT_ID"],
		envVars["GOOGLE_OAUTH_CLIENT_SECRET"],
		envVars["OAUTH_REDIRECT_URL"],
	}
)

func getEnvVariables() map[string]string {
	vars := make(map[string]string)

	var serverDirPath = utils.GetExecutableDirPath()
	var envPath = filepath.Join(serverDirPath, "../.env")

	_ = godotenv.Load(envPath)

	for _, requiredVar := range requiredVars {
		value := os.Getenv(requiredVar)

		if value == "" {
			value = defaultVars[requiredVar]
		}

		if value == "" {
			fmt.Printf("Missing env variable %v\n", requiredVar)
			os.Exit(1)
		}

		vars[requiredVar] = value
	}

	return vars
}
