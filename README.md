# Description
A realtime web application to help development teams vote for user story points estimation during sprint planning meetings. Voting options are currently fixed, but in a future release, they will be chosen as a setting for each team.

This repository contains the server-side code. The front-end code is here: https://github.com/ibrahimfarhan/voting-app-webapp. It also contains an example `docker-compose` file to get the stack up and running quickly.
# Running
The application can run by either:
### Installing dependencies on your machine:
1. Download and install Go: [https://golang.org](https://golang.org).
2. Download and install MongoDB: [https://www.mongodb.com/try/download/community](https://www.mongodb.com/try/download/community).
3. Create a .env file and assign values to the required env variables as demonstrated in the example.env file. Environment variables will be explained in the next section.
4. In the project directory, run `go run main/main.go`

### Or using Docker
This method runs the whole stack, server, webapp and DB. Therefore, the `voting-app-webapp` folder must be in the same folder as `voting-app-server`, i.e, on the same level, in order for the docker-compose file to be able to build the webapp image.
1. Download and install Docker if not existent on your local machine.
2. Clone this repository and the `voting-app-webapp` repository and put both in the same folder.
3. Edit the `environments` section of the `voting-app-server` service in the `example.docker-compose.yaml` file as shown in the next section.
4. Run `docker-compose -f example.docker-compose.yaml up`.

# Environment Variables
The `voting-app-server` envs:
- `PORT` The port on which the server listens.
- `Env` The running environment. Accepts `development` or `production`.
- `DB_TYPE` The only supported option currently is `mongo`.
- `DB_CONNECTION_STRING`
- `DB_NAME`
- `ALLOWED_ORIGINS` The URLs of the front-end applications allowed to access this server in a comma separated format.
- `SERVE_STATIC_FILES` Accepts `true` or `false`. Default is `false`. Determines whether this server should act as a static files server. In this case the repository must contain a folder named `client` containing the static files.
- `GOOGLE_OAUTH_CLIENT_ID` and `GOOGLE_OAUTH_CLIENT_SECRET` are required for enabling logging-in with Google.
- `OAUTH_REDIRECT_URL` The URL that the OAuth provider redirects the user to after logging in. Must be in the following structure: `{{apiServerBaseURL}}/api/v1/user/oauth/provider/callback`.

The `voting-app-webapp` envs:
- `REACT_APP_API_SERVER_URL`
- `REACT_APP_WEBSOCKET_URL`

The provided `example.docker-compose.yaml` already contains default values, so you can edit only whichever suits your settings.

