package realtime

import "github.com/ibrahimfarhan/voting-app/voting-app-server/logger"

type VotingHub struct {
	TeamSessions   map[string]*teamSession
	NewSessions    chan *teamSession
	EndingSessions chan *teamSession
}

func NewVotingHub() *VotingHub {
	return &VotingHub{
		TeamSessions:   map[string]*teamSession{},
		NewSessions:    make(chan *teamSession),
		EndingSessions: make(chan *teamSession),
	}
}

func (h *VotingHub) Run() {
	logger.Info("Voting hub is started")

	for {
		select {
		case s := <-h.NewSessions:
			h.TeamSessions[s.teamID] = s
			logger.Info("Added a new voting session for " + s.teamID)

		case s := <-h.EndingSessions:
			if _, ok := h.TeamSessions[s.teamID]; ok {
				delete(h.TeamSessions, s.teamID)
				logger.Info("Deleted voting session of" + s.teamID)
			}
		}
	}
}
