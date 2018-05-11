package memory

import (
	"errors"

	"github.com/tomocy/goron/session"
)

var sessions = make(map[string]session.Session)

type memory struct {
}

func New() *memory {
	return &memory{}
}

func (m *memory) InitSession(sessionID string) session.Session {
	dat := make(map[string]interface{})
	sessions[sessionID] = session.New(sessionID, dat)

	return sessions[sessionID]
}

func (m *memory) GetSession(sessionID string) (session.Session, error) {
	session, ok := sessions[sessionID]
	if !ok {
		return nil, errors.New("Session not found")
	}

	return session, nil
}

func (m *memory) DestroySession(sessionID string) error {
	if _, ok := sessions[sessionID]; ok {
		delete(sessions, sessionID)
		return nil
	}

	return errors.New("Session not found: could not delete the session")
}
