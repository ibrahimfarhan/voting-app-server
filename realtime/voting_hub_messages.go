package realtime

import (
	"encoding/json"
)

/**
*  Hub actions have three types:
*    Incoming: sent by the client to the server.
*    Outgoing: sent by the server to the client.
*    Two-way: received first by the server and then sent to the client.
*	 The client needs to handle outgoing and two-way actions.
**/
var hubActions = struct {
	// Incoming Actions:
	closeSession string
	submitVote   string

	// Outgoing Actions:
	showNewMember       string
	removeLeavingMember string
	showCurrentMembers  string
	sessionCreated      string
	showVote            string

	// Two-way
	displaySubject       string
	toggleResultsDisplay string
	reset                string
	changeModerator      string
}{
	"closeSession",
	"submitVote",
	"showNewMember",
	"removeLeavingMember",
	"showCurrentMembers",
	"sessionCreated",
	"showVote",
	"displaySubject",
	"toggleResultsDisplay",
	"reset",
	"changeModerator",
}

var indicativeVotingOptions = struct {
	voted   int8
	unknown int8
}{
	-1,
	-2,
}

type webSocketMessage struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data,omitempty"`
	sender *sessionMember
}

type memberInMessage struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type showCurrentMembersMessage struct {
	SessionModerator memberInMessage   `json:"sessionModerator"`
	CurrentMembers   []memberInMessage `json:"currentMembers"`
}

func validateMessage(msg []byte, member *sessionMember) (bool, *webSocketMessage) {
	webSocketMsg := new(webSocketMessage)
	err := json.Unmarshal(msg, webSocketMsg)
	if err != nil || webSocketMsg.Action == "" {
		return false, nil
	}

	webSocketMsg.sender = member

	memberIsSessionModerator := member.id == member.currentSession.moderator.id

	if webSocketMsg.Action == hubActions.submitVote {
		_, ok := webSocketMsg.Data.(float64)
		if !ok {
			return false, nil
		}

	} else {
		if !memberIsSessionModerator {
			return false, nil
		}
	}

	return true, webSocketMsg
}
