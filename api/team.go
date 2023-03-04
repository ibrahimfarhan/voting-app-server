package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/config"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/logger"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/store/storeutils"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"
)

func (api *Api) registerTeamRoutes() {
	tr := api.routes.team
	tr.Handle("/all-public", api.handleAuthRequired(getPublicNonJoinedTeams)).Methods("GET", "OPTIONS")
	tr.Handle("/create", api.handleAuthRequired(createTeam)).Methods("POST", "OPTIONS")
	tr.Handle("/join/{tokenOrTeamID}", api.handleAuthRequired(joinTeam)).Methods("PATCH", "OPTIONS")
	tr.Handle("/leave", api.handleAuthRequired(leaveTeam)).Methods("PATCH", "OPTIONS")
	tr.Handle("/toggle-role", api.handleAuthRequired(toggleRole)).Methods("PATCH", "OPTIONS")
	tr.Handle("/add-member", api.handleAuthRequired(addMember)).Methods("PATCH", "OPTIONS")
	tr.Handle("/remove-member", api.handleAuthRequired(removeMember)).Methods("PATCH", "OPTIONS")
	tr.Handle("/{id}", api.handleAuthRequired(updateTeam)).Methods("PATCH", "OPTIONS")
	tr.Handle("/{id}", api.handleAuthRequired(deleteTeam)).Methods("DELETE", "OPTIONS")
	tr.Handle("/{id}/invite", api.handleAuthRequired(getInvitationToken)).Methods("POST", "OPTIONS")
	tr.Handle("/{id}", api.handleAuthRequired(getTeam)).Methods("GET", "OPTIONS")
}

func getPublicNonJoinedTeams(c *apiContext, w http.ResponseWriter, req *http.Request) {
	teams, err := c.app.TeamStore.GetPublicNonJoined(c.user.ID)
	if err != nil {
		logger.Error("teamAPI.getPublicNonJoined", err)
		sendErrorResponse("Error while getting teams", http.StatusInternalServerError, w)
		return
	}

	sendJSONResponse(teams, http.StatusOK, w)
}

func getTeam(c *apiContext, w http.ResponseWriter, req *http.Request) {
	team, _, err := checkUserIsInTeam(mux.Vars(req)["id"], c.user.ID, c)
	if team != nil && !team.IsPublic && err != nil {
		sendErrorResponse(models.Forbidden, http.StatusForbidden, w)
		return
	}

	sendJSONResponse(team, http.StatusOK, w)
}

func createTeam(c *apiContext, w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	isValid, t := models.ValidateTeamData(reqBody)
	if !isValid {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	teamsCount, err := c.app.TeamStore.GetCountByUserID(c.user.ID)
	if err != nil {
		sendErrorResponse(models.SomethingWentWrong, http.StatusInternalServerError, w)
		return
	}

	if teamsCount >= int64(config.Env.MaxTeamsCountPerUser) {
		sendErrorResponse("Maximum number of teams reached", http.StatusBadRequest, w)
		return
	}

	savedTeam, err := c.app.TeamStore.GetOne(storeutils.QueryOptions{
		Conditions: storeutils.M{"Name": t.Name},
	})
	if savedTeam != nil && err == nil {
		sendErrorResponse("Team name already exists", http.StatusBadRequest, w)
		return
	}

	t.AdminIDs = []string{c.user.ID}
	team, err := c.app.TeamStore.Create(t)
	if err != nil {
		sendErrorResponse(models.SomethingWentWrong, http.StatusInternalServerError, w)
		return
	}

	sendJSONResponse(team, http.StatusOK, w)
}

func updateTeam(c *apiContext, w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	teamID := mux.Vars(req)["id"]

	isValid, editedTeam := models.ValidateTeamData(reqBody)
	if !isValid || !models.IsValidID(teamID) {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	editedTeam.ID = teamID

	_, err := checkUserIsTeamAdmin(editedTeam.ID, c)
	if err != nil {
		sendErrorResponse(models.Forbidden, http.StatusForbidden, w)
		return
	}

	team, err := c.app.TeamStore.Update(editedTeam)
	if err != nil {
		sendErrorResponse("Team was not modified", http.StatusBadRequest, w)
		return
	}

	sendJSONResponse(team, http.StatusOK, w)
}

func deleteTeam(c *apiContext, w http.ResponseWriter, req *http.Request) {
	teamID := mux.Vars(req)["id"]

	if !models.IsValidID(teamID) {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	_, err := checkUserIsTeamAdmin(teamID, c)
	if err != nil {
		logger.Error("deleteTeam checkUserIsTeamAdmin", err)
		sendErrorResponse(models.Forbidden, http.StatusForbidden, w)
		return
	}

	err = c.app.TeamStore.DeleteByID(teamID)
	if err != nil {
		sendErrorResponse("Error while deleting team", http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getInvitationToken(c *apiContext, w http.ResponseWriter, req *http.Request) {
	teamID := mux.Vars(req)["id"]

	if !models.IsValidID(teamID) {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	team, _, err := checkUserIsInTeam(teamID, c.user.ID, c)
	if err != nil {
		sendErrorResponse(models.Forbidden, http.StatusForbidden, w)
		return
	}

	existentToken, err := c.app.TokenStore.GetOne(storeutils.QueryOptions{
		Conditions: storeutils.M{"OwnerID": teamID},
	})
	if err == nil && existentToken != nil && existentToken.ExpiresAt.After(time.Now()) {
		sendJSONResponse(map[string]string{"token": existentToken.ID}, http.StatusOK, w)
		return
	}

	token := &models.Token{
		ID:        utils.NewLongID(),
		Type:      models.InviteToken,
		ExpiresAt: time.Now().Add(models.InviteTokenExpiryDuration),
		OwnerID:   team.ID,
	}

	t, err := c.app.TokenStore.Create(token)
	if err != nil {
		sendErrorResponse("Could not create invitation link", http.StatusBadRequest, w)
		return
	}

	sendJSONResponse(map[string]string{"token": t.ID}, http.StatusOK, w)
}

func joinTeam(c *apiContext, w http.ResponseWriter, req *http.Request) {
	tokenOrTeamID := mux.Vars(req)["tokenOrTeamID"]
	teamID := ""
	isJoiningPublicTeam := len(tokenOrTeamID) != models.TokenIDLength

	if !isJoiningPublicTeam {
		t, err := c.app.TokenStore.GetByID(tokenOrTeamID)
		if err != nil {
			sendErrorResponse("Invalid invitation link", http.StatusBadRequest, w)
			return
		}

		if t.Type != models.InviteToken || t.ExpiresAt.Before(time.Now()) {
			sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
			return
		}

		teamID = t.OwnerID
	} else {
		teamID = tokenOrTeamID
	}

	team, _, err := checkUserIsInTeam(teamID, c.user.ID, c)
	if team == nil {
		sendErrorResponse("Team not found", http.StatusNotFound, w)
		return
	}
	if !team.IsPublic && isJoiningPublicTeam {
		sendErrorResponse("Forbidden", http.StatusForbidden, w)
		return
	}
	if err == nil {
		sendErrorResponse("User already exists in team", http.StatusBadRequest, w)
		return
	}

	err = c.app.TeamStore.AddMemberToTeam(teamID, c.user.ID)
	if err != nil {
		sendErrorResponse("User was not added to the team", http.StatusBadRequest, w)
		return
	}

	sendJSONResponse(team, http.StatusOK, w)
}

func leaveTeam(c *apiContext, w http.ResponseWriter, req *http.Request) {
	teamID := utils.MapFromBody(req.Body)["id"].(string)

	if !models.IsValidID(teamID) {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	team, isAdmin, err := checkUserIsInTeam(teamID, c.user.ID, c)
	if err != nil {
		sendErrorResponse("User is not a member in the team", http.StatusForbidden, w)
		return
	}

	if isAdmin {
		if len(team.Admins) == 1 && len(team.Members) == 0 {
			sendErrorResponse("Cannot leave the team because it contains no other members. Delete the team instead.", http.StatusBadRequest, w)
			return
		}
		err := c.app.TeamStore.RemoveAdminFromTeam(team, c.user.ID)
		if err != nil {
			sendErrorResponse("Could not remove admin from the team", http.StatusBadRequest, w)
			return
		}
	} else {
		err = c.app.TeamStore.RemoveMemberFromTeam(teamID, c.user.ID)
		if err != nil {
			sendErrorResponse("Could not remove user from the team", http.StatusBadRequest, w)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func toggleRole(c *apiContext, w http.ResponseWriter, req *http.Request) {
	bodyData := utils.MapFromBody(req.Body)

	teamID, ok := bodyData["teamId"].(string)
	userID, ok := bodyData["userId"].(string)
	if !ok {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	if !models.IsValidID(teamID) {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	t, err := checkUserIsTeamAdmin(teamID, c)
	if err != nil {
		sendErrorResponse(models.Forbidden, http.StatusForbidden, w)
		return
	}

	isAdmin := false
	for _, a := range t.Admins {
		if a.ID == userID {
			isAdmin = true
			break
		}
	}

	isMember := false
	for _, a := range t.Members {
		if a.ID == userID {
			isMember = true
			break
		}
	}

	if !isAdmin && !isMember {
		sendErrorResponse("User is not a member in the team", http.StatusBadRequest, w)
		return
	}

	err = c.app.TeamStore.ToggleRole(t, userID, isAdmin)
	if err != nil {
		sendErrorResponse("Error while changing user role", http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func addMember(c *apiContext, w http.ResponseWriter, req *http.Request) {
}

func removeMember(c *apiContext, w http.ResponseWriter, req *http.Request) {
	bodyData := utils.MapFromBody(req.Body)

	teamID, ok := bodyData["teamId"].(string)
	userID, ok := bodyData["userId"].(string)
	if !ok {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	if !models.IsValidID(teamID) {
		sendErrorResponse(models.ValidationError, http.StatusBadRequest, w)
		return
	}

	t, err := checkUserIsTeamAdmin(teamID, c)
	if err != nil {
		sendErrorResponse(models.Forbidden, http.StatusForbidden, w)
		return
	}

	isAdmin := false
	for _, a := range t.Admins {
		if a.ID == userID {
			isAdmin = true
			break
		}
	}

	isMember := false
	for _, a := range t.Members {
		if a.ID == userID {
			isMember = true
			break
		}
	}

	if !isAdmin && !isMember {
		sendErrorResponse("User is not a member in the team", http.StatusBadRequest, w)
		return
	}

	if isAdmin {
		if len(t.Admins) == 1 && len(t.Members) == 0 {
			sendErrorResponse("Error while removing user", http.StatusBadRequest, w)
			return
		}
		err := c.app.TeamStore.RemoveAdminFromTeam(t, userID)
		if err != nil {
			sendErrorResponse("Error while removing user", http.StatusInternalServerError, w)
			return
		}
	} else {
		err = c.app.TeamStore.RemoveMemberFromTeam(teamID, userID)
		if err != nil {
			sendErrorResponse("Error while removing user", http.StatusInternalServerError, w)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

/*------------------------------------- Helpers -------------------------------------*/

func checkUserIsInTeam(teamID, userID string, c *apiContext) (team *models.Team, isAdmin bool, err error) {
	team, err = c.app.TeamStore.GetByIDWithUsers(teamID)
	if err != nil {
		return nil, false, err
	}

	for _, u := range team.Members {
		if u.ID == userID {
			return team, false, nil
		}
	}

	for _, a := range team.Admins {
		if a.ID == userID {
			return team, true, nil
		}
	}

	return team, false, errors.New(models.Forbidden)
}

func checkUserIsTeamAdmin(teamID string, c *apiContext) (*models.Team, error) {
	team, isAdmin, err := checkUserIsInTeam(teamID, c.user.ID, c)

	if err != nil {
		return nil, err
	}

	if !isAdmin {
		return nil, errors.New(models.Forbidden)
	}

	return team, nil
}
