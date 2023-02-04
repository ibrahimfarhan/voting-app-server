package realtime

import (
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/logger"
	"github.com/ibrahimfarhan/voting-app/voting-app-server/models"
)

const (
	writeWait          = 10 * time.Second
	pongWait           = 60 * time.Second
	pingPeriod         = (pongWait * 9) / 10
	maxMessageSize     = 512
	outboundMsgsBuffer = 3
)

type sessionMember struct {
	id                 string
	username           string
	isSessionModerator bool
	outboundMsgs       chan []byte
	currentSession     *teamSession
	conn               *websocket.Conn
	hub                *VotingHub
	left               bool
	mu                 *sync.Mutex
}

func newSessionMember(user *models.User, currentSession *teamSession, conn *websocket.Conn, hub *VotingHub) *sessionMember {
	username := user.Username
	if user.Username == "" {
		username = user.Name
	}
	return &sessionMember{
		id:             user.ID,
		username:       username,
		outboundMsgs:   make(chan []byte, outboundMsgsBuffer),
		currentSession: currentSession,
		conn:           conn,
		hub:            hub,
		mu:             &sync.Mutex{},
	}
}

func AddSessionMember(user *models.User, currentSession *teamSession, conn *websocket.Conn, hub *VotingHub) {
	// Make sure the user has not already joined the session as admin or as member.
	_, ok := currentSession.members[user.Username]
	if user.ID == currentSession.moderator.id || ok {
		return
	}

	m := newSessionMember(user, currentSession, conn, hub)

	m.currentSession.newMembers <- m
}

func (m *sessionMember) read() {
	logger.Info("Started reading web socket messages for user " + m.username)

	m.conn.SetReadLimit(maxMessageSize)
	m.conn.SetReadDeadline(time.Now().Add(pongWait))
	m.conn.SetPongHandler(func(string) error { m.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		if m.left {
			return
		}

		_, message, err := m.conn.ReadMessage()
		if err != nil {
			logger.Error("sessionMember.read for "+m.username, err)
			m.leave()
			return
		}

		isValid, webSocketMsg := validateMessage(message, m)
		if isValid {
			m.currentSession.messages <- webSocketMsg
		}
	}
}

func (m *sessionMember) write() {
	logger.Info("Started writing web socket messages for user " + m.username)

	ticker := time.NewTicker(pingPeriod)

	for {
		if m.left {
			return
		}

		select {
		case message, ok := <-m.outboundMsgs:
			m.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				m.conn.WriteMessage(websocket.CloseMessage, []byte{})
				m.handleWriteError(errors.New("outboundMsgs is closed for user " + m.username))
				return
			}

			if err := m.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				m.handleWriteError(err)
				return
			}

			n := len(m.outboundMsgs)
			for i := 0; i < n; i++ {
				if err := m.conn.WriteMessage(websocket.TextMessage, <-m.outboundMsgs); err != nil {
					m.handleWriteError(err)
					return
				}
			}

		case <-ticker.C:
			m.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := m.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				m.handleWriteError(err)
				return
			}
		}
	}
}

func (m *sessionMember) handleWriteError(err error) {
	logger.Error("sessionMember.write for "+m.username, err)
	m.leave()
}

func (m *sessionMember) leave() {
	m.mu.Lock()

	if m.left {
		m.mu.Unlock()
		return
	}

	m.left = true

	m.mu.Unlock()

	logger.Info("User " + m.username + " is leaving session")

	if m == m.currentSession.moderator {
		m.currentSession.stop()
	} else {
		if m.currentSession.isRunning {
			m.currentSession.leavingMembers <- m
		}
	}

	close(m.outboundMsgs)
	m.conn.Close()
}
