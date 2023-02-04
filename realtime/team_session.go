package realtime

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/logger"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
)

type teamSession struct {
	teamID          string
	moderator       *sessionMember
	members         map[string]*sessionMember
	votes           map[string]int8
	currentSubject  string
	messages        chan *webSocketMessage
	newMembers      chan *sessionMember
	leavingMembers  chan *sessionMember
	hub             *VotingHub
	quit            chan bool
	isRunning       bool
	resultsAreShown bool
}

func AddTeamSession(teamID string, user *models.User, conn *websocket.Conn, hub *VotingHub) {
	s := &teamSession{
		teamID:          teamID,
		members:         map[string]*sessionMember{},
		votes:           map[string]int8{},
		currentSubject:  "",
		messages:        make(chan *webSocketMessage),
		newMembers:      make(chan *sessionMember),
		leavingMembers:  make(chan *sessionMember),
		hub:             hub,
		quit:            make(chan bool),
		isRunning:       true,
		resultsAreShown: false,
	}

	a := newSessionMember(user, s, conn, hub)
	s.moderator = a

	s.hub.NewSessions <- s

	go s.start()

	go a.read()
	go a.write()

	msg, err := json.Marshal(webSocketMessage{Action: hubActions.sessionCreated})
	if err == nil {
		s.moderator.outboundMsgs <- msg
	}
}

func (s *teamSession) UserExistsInSession(user *models.User) bool {
	_, ok := s.members[user.ID]
	if user.ID == s.moderator.id || ok {
		return true
	}

	return false
}

func (s *teamSession) start() {
	logger.Info("Starting voting session for team " + s.teamID)

	for {
		select {
		case m := <-s.newMembers:
			s.handleNewMember(m)

		case m := <-s.leavingMembers:
			s.handleLeavingMember(m)

		case a := <-s.messages:
			data, err := json.Marshal(a)
			if err != nil {
				break
			}

			s.handleMessage(a.Action, data, a)

		case <-s.quit:
			logger.Info("Voting session is stopped for team " + s.teamID)
			return
		}
	}
}

func (s *teamSession) stop() {
	logger.Info("Stopping voting session for team " + s.teamID)
	s.hub.EndingSessions <- s
	s.quit <- true
	s.isRunning = false

	for _, m := range s.members {
		go m.leave()
	}

	close(s.messages)
	close(s.newMembers)
	close(s.leavingMembers)
	close(s.quit)
}

func (s *teamSession) handleMessage(action string, msg []byte, webSocketMsg *webSocketMessage) {
	switch action {
	case hubActions.closeSession:
		go s.moderator.leave()

	case hubActions.displaySubject:
		s.handleDisplaySubject(msg, webSocketMsg)

	case hubActions.reset:
		s.reset(msg)

	case hubActions.submitVote:
		s.handleSubmitVote(webSocketMsg)

	case hubActions.toggleResultsDisplay:
		s.handleToggleResultsDisplay()

	case hubActions.changeModerator:
		s.handleChangeModerator(msg, webSocketMsg)
	}
}

/*------------------------------------- Methods to handle hub actions -------------------------------------*/

func (s *teamSession) handleDisplaySubject(msg []byte, webSocketMsg *webSocketMessage) {
	s.currentSubject = webSocketMsg.Data.(string)
	s.broadcast(msg)
}

func (s *teamSession) reset(msg []byte) {
	s.votes = map[string]int8{}
	s.resultsAreShown = false
	s.broadcast(msg)
}

func (s *teamSession) handleSubmitVote(webSocketMsg *webSocketMessage) {
	senderId := webSocketMsg.sender.id
	points := int8(webSocketMsg.Data.(float64))

	vote := map[string]int8{senderId: points}
	msg, err := json.Marshal(webSocketMessage{Action: hubActions.showVote, Data: vote})
	if err != nil {
		return
	}

	var msgToSendToOthers []byte

	if s.resultsAreShown {
		msgToSendToOthers = msg
	} else {
		hiddenVote := map[string]int8{}
		s.votes[webSocketMsg.sender.id] = points
		hiddenVote[webSocketMsg.sender.id] = indicativeVotingOptions.voted

		hiddenVotingMsg, err := json.Marshal(webSocketMessage{
			Action: hubActions.showVote, Data: hiddenVote,
		})
		if err != nil {
			return
		}

		msgToSendToOthers = hiddenVotingMsg
	}

	if senderId == s.moderator.id {
		s.moderator.outboundMsgs <- msg
		s.sendMessageToMembers(msgToSendToOthers)
	} else {
		s.members[senderId].outboundMsgs <- msg
		s.sendToAllExcept(msgToSendToOthers, webSocketMsg.sender)
	}
}

func (s *teamSession) handleNewMember(m *sessionMember) {
	s.members[m.id] = m
	go m.read()
	go m.write()

	// Send new member to all connected users.
	msg := &webSocketMessage{Action: hubActions.showNewMember, Data: memberInMessage{
		ID:       m.id,
		Username: m.username,
	}}
	broadcastMsg, err := json.Marshal(msg)
	if err == nil {
		s.broadcast(broadcastMsg)
	}

	// Send current members and current subject to the new member.
	currentMembers := []memberInMessage{}
	sessionModerator := memberInMessage{
		ID:       s.moderator.id,
		Username: s.moderator.username,
	}

	for _, member := range s.members {
		currentMembers = append(currentMembers, memberInMessage{
			ID:       member.id,
			Username: member.username,
		})
	}

	showCurrentMembersMsg := webSocketMessage{
		Action: hubActions.showCurrentMembers,
		Data:   showCurrentMembersMessage{SessionModerator: sessionModerator, CurrentMembers: currentMembers},
	}

	displaySubjectMsg := webSocketMessage{Action: hubActions.displaySubject, Data: s.currentSubject}

	webSocketMsgs := []webSocketMessage{showCurrentMembersMsg, displaySubjectMsg}

	votesToSend := map[string]int8{}

	votesMsg := webSocketMessage{}

	if s.resultsAreShown {
		votesMsg.Action = hubActions.toggleResultsDisplay
		votesMsg.Data = map[string]interface{}{"votes": s.votes, "show": true}
	} else {
		votesMsg.Action = hubActions.showVote
		for k, v := range s.votes {
			if v != 0 {
				votesToSend[k] = indicativeVotingOptions.voted
			}
		}
		votesMsg.Data = votesToSend
	}

	webSocketMsgs = append(webSocketMsgs, votesMsg)

	msgs, err := json.Marshal(webSocketMsgs)
	if err == nil {
		m.outboundMsgs <- msgs
	}
}

func (s *teamSession) handleLeavingMember(m *sessionMember) {
	if _, ok := s.members[m.id]; ok {
		delete(s.members, m.id)

		if _, ok := s.votes[m.id]; ok {
			delete(s.votes, m.id)
		}

		msg := &webSocketMessage{Action: hubActions.removeLeavingMember, Data: memberInMessage{ID: m.id, Username: m.username}}
		broadcastMsg, err := json.Marshal(msg)
		if err == nil {
			s.broadcast(broadcastMsg)
		}
	}
}

func (s *teamSession) handleToggleResultsDisplay() {
	msg := webSocketMessage{Action: hubActions.toggleResultsDisplay}

	if !s.resultsAreShown {
		msg.Data = map[string]interface{}{"votes": s.votes, "show": true}
	}

	m, err := json.Marshal(msg)
	if err != nil {
		return
	}

	s.resultsAreShown = !s.resultsAreShown

	s.broadcast(m)
}

func (s *teamSession) handleChangeModerator(msg []byte, webSocketMessage *webSocketMessage) {
	newModeratorId := webSocketMessage.Data.(string)
	s.members[s.moderator.id] = s.moderator
	newModerator, ok := s.members[newModeratorId]
	if !ok {
		return
	}
	delete(s.members, newModeratorId)
	s.moderator = newModerator
	s.reset(msg)
}

/*------------------------------------- Methods to send messages to connected users -------------------------------------*/

func (s *teamSession) sendMessageToMembers(msg []byte) {
	for _, m := range s.members {
		go func(m *sessionMember) {
			m.outboundMsgs <- msg
		}(m)
	}
}

func (s *teamSession) broadcast(msg []byte) {
	go func(admin *sessionMember) {
		admin.outboundMsgs <- msg
	}(s.moderator)

	s.sendMessageToMembers(msg)
}

func (s *teamSession) sendToAllExcept(msg []byte, excludedMember *sessionMember) {
	go func(admin *sessionMember) {
		admin.outboundMsgs <- msg
	}(s.moderator)

	for _, m := range s.members {
		if m.id != excludedMember.id {
			go func(m *sessionMember) {
				m.outboundMsgs <- msg
			}(m)
		}
	}
}
