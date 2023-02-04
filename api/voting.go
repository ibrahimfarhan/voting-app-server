package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/logger"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/realtime"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (api *Api) registerVotingRoutes() {
	tr := api.routes.voting
	tr.Handle("/{teamId}", api.handleAuthRequired(startSession)).Methods("GET", "OPTIONS")
}

func startSession(c *apiContext, w http.ResponseWriter, req *http.Request) {
	teamID := mux.Vars(req)["teamId"]
	if !models.IsValidID(teamID) {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	_, _, err := checkUserIsInTeam(teamID, c.user.ID, c)
	if err != nil {
		sendErrorResponse(models.Forbidden, http.StatusForbidden, w)
		return
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	teamSession, ok := c.app.VotingHub.TeamSessions[teamID]
	if ok && teamSession.UserExistsInSession(c.user) {
		sendErrorResponse("User has already joined the running session", http.StatusBadRequest, w)
		return
	}

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error("startSession Upgrade", err)
		return
	}

	if ok {
		realtime.AddSessionMember(c.user, teamSession, conn, c.app.VotingHub)
	} else {
		realtime.AddTeamSession(teamID, c.user, conn, c.app.VotingHub)
	}
}
