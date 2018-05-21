package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/tomocy/goron/session"
	"github.com/tomocy/goron/settings"
)

var sessions = make(map[string]session.Session)

type memory struct {
	mu sync.Mutex
}

func New() *memory {
	return &memory{}
}

func (m *memory) InitSession(sessionID string) session.Session {
	m.mu.Lock()
	defer m.mu.Unlock()

	dat := make(map[string]string)
	sessions[sessionID] = session.New(sessionID, time.Now().Add(settings.Session.ExpiresIn), dat)

	return sessions[sessionID]
}

func (m *memory) GetSession(sessionID string) (session.Session, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, ok := sessions[sessionID]
	if !ok {
		return nil, errors.New("Session not found")
	}

	return session, nil
}

func (m *memory) SetSession(session session.Session) {
	m.mu.Lock()
	defer m.mu.Unlock()

	sessions[session.ID()] = session
}

func (m *memory) DeleteSession(sessionID string) {

}

func (m *memory) DeleteExpiredSessions() {

}
