package interfaces

import "github.com/ibrahimfarhan/voting-app/voting-app-server/models"

type TeamStore interface {
	entityStore[models.Team]
	GetPublicNonJoined(userID string) ([]*models.Team, error)
	GetByIDWithUsers(id string) (*models.Team, error)
	GetByUserID(userID string, includeUsers bool) ([]*models.Team, error)
	GetCountByUserID(userID string) (int64, error)
	AddMemberToTeam(teamID, memberID string) error
	AddAdminToTeam(team *models.Team, memberID string) error
	RemoveMemberFromTeam(teamID, memberID string) error
	RemoveAdminFromTeam(team *models.Team, adminID string) error
	ToggleRole(team *models.Team, userID string, isAdmin bool) error
}
