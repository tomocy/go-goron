package memory

import (
	"errors"
	"sync"

	"github.com/tomocy/goron/session"
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

	dat := make(map[string]interface{})
	sessions[sessionID] = session.New(sessionID, dat)

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
