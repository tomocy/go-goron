package manager

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/tomocy/goron/session"
	"github.com/tomocy/goron/session/cookie"
	"github.com/tomocy/goron/session/storages"
)

type Manager interface {
	GetSession(w http.ResponseWriter, r *http.Request) session.Session
	SetSession(session session.Session)

	generateSessionID() string
}

type manager struct {
	storage storages.Storage
}

func New(storageName string) (Manager, error) {
	storage, err := storages.Get(storageName)
	if err != nil {
		return nil, fmt.Errorf("Storage not found: %s", storageName)
	}

	return &manager{storage: storage}, nil
}

func (m *manager) GetSession(w http.ResponseWriter, r *http.Request) session.Session {
	sessionID, err := cookie.GetSessionID(r)
	if err != nil {
		// No session id in client
		// Start new session with new sesion id
		sessionID = m.generateSessionID()
		cookie.SetSessionID(w, sessionID)

		return m.storage.InitSession(sessionID)
	}

	session, err := m.storage.GetSession(sessionID)
	if err != nil {
		// No session in server while client has session id
		// Delete session id in client and start new session with new session id
		cookie.DestroySessionID(w)

		sessionID = m.generateSessionID()
		cookie.SetSessionID(w, sessionID)

		return m.storage.InitSession(sessionID)
	}

	return session
}

func (m *manager) SetSession(session session.Session) {
	m.storage.SetSession(session)
}
func (m *manager) generateSessionID() string {
	return uuid.New().String()
}
