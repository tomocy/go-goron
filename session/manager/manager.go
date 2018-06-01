package manager

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tomocy/goron/session"
	"github.com/tomocy/goron/session/cookie"
	"github.com/tomocy/goron/session/storages"
	"github.com/tomocy/goron/settings"
)

type Manager struct {
	storage             storages.Storage
	probOfDelete        int
	probOfDeleteDivisor int
}

var m *Manager

func GetReady() *Manager {
	return m
}

func (m *Manager) GetSession(w http.ResponseWriter, r *http.Request) *session.Session {
	sessionID, err := cookie.GetSessionID(r)
	if err != nil {
		// No session id in client
		// Start new session with new sesion id
		return m.recreateSession(w)
	}

	session, err := m.storage.GetSession(sessionID)
	if err != nil {
		// No session in server while client has session id
		// Start new session
		return m.recreateSession(w)
	}

	if session.DoesExpire() {
		// When session expires
		// Delete session in serve and start new session
		return m.recreateSession(w)
	}

	return session
}

func (m *Manager) SetSession(session *session.Session) {
	m.storage.SetSession(session)
}

func (m *Manager) deleteExpiredSessions() {
	t := time.NewTicker(1 * time.Minute)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			if m.doesDelete() {
				m.storage.DeleteExpiredSessions()
			}
		}
	}
}

func (m *Manager) recreateSession(w http.ResponseWriter) *session.Session {
	sessionID := generateSessionID()
	cookie.SetSessionID(w, sessionID)

	return m.storage.InitSession(sessionID)
}

func (m *Manager) doesDelete() bool {
	if m.probOfDelete <= 0 {
		return false
	}
	if m.probOfDeleteDivisor <= m.probOfDelete {
		return true
	}

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(m.probOfDeleteDivisor)+1 <= m.probOfDelete
}

func generateSessionID() string {
	return uuid.New().String()
}

func new(storageName string) (*Manager, error) {
	storage, err := storages.Get(storageName)
	if err != nil {
		return nil, fmt.Errorf("Storage not found: %s", storageName)
	}

	m := &Manager{
		storage:             storage,
		probOfDelete:        settings.SessionManager.ProbOfDelete(),
		probOfDeleteDivisor: settings.SessionManager.ProbOfDeleteDivisor(),
	}

	return m, nil
}

func init() {
	var err error
	m, err = new(settings.Session.Storage())
	if err != nil {
		panic(err)
	}

	go m.deleteExpiredSessions()
}
