package models

import (
	"encoding/json"
	"time"
)

const (
	maxTeamNameLength = 100
	minTeamNameLength = 3
)

type Team struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Admins    []*User   `json:"admins,omitempty" bson:"admins,omitempty"`
	Members   []*User   `json:"members,omitempty" bson:"members,omitempty"`
	AdminIDs  []string  `json:"adminIds,omitempty" bson:"adminIds,omitempty"`
	MemberIDs []string  `json:"memberIds,omitempty" bson:"memberIds,omitempty"`
	IsPublic  bool      `json:"isPublic,omitempty" bson:"isPublic,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}

func ValidateTeamData(teamData []byte) (bool, *Team) {
	t := new(Team)

	err := json.Unmarshal(teamData, t)
	if err != nil {
		return false, nil
	}

	return len(t.Name) <= maxTeamNameLength && len(t.Name) >= minTeamNameLength && nameRegexp.MatchString(t.Name), t
}
