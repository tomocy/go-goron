package manager

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/tomocy/goron/log"
	"github.com/tomocy/goron/session"
	"github.com/tomocy/goron/session/cookie"
	"github.com/tomocy/goron/session/storages"
	"github.com/tomocy/goron/settings"
)

type Manager interface {
	GetSession(w http.ResponseWriter, r *http.Request) session.Session
	SetSession(session session.Session)
	DeleteExpiredSessions()

	generateSessionID() string
}

type manager struct {
	storage             storages.Storage
	probOfDelete        int
	probOfDeleteDivisor int
}

func New(storageName string) (Manager, error) {
	storage, err := storages.Get(storageName)
	if err != nil {
		return nil, fmt.Errorf("Storage not found: %s", storageName)
	}

	m := &manager{
		storage:             storage,
		probOfDelete:        settings.SessionManager.ProbOfDelete,
		probOfDeleteDivisor: settings.SessionManager.ProbOfDeleteDivisor,
	}

	return m, nil
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
		// Start new session
		sessionID = m.generateSessionID()
		cookie.SetSessionID(w, sessionID)

		return m.storage.InitSession(sessionID)
	}

	if session.DoesExpire() {
		// When session expires
		// Delete session in serve and start new session
		m.storage.DeleteSession(sessionID)
		sessionID = m.generateSessionID()
		cookie.SetSessionID(w, sessionID)

		return m.storage.InitSession(sessionID)
	}

	return session
}

func (m *manager) SetSession(session session.Session) {
	m.storage.SetSession(session)
}

func (m *manager) DeleteExpiredSessions() {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			if m.doesDelete() {
				log.Debug("true")
				m.storage.DeleteExpiredSessions()
			} else {
				log.Debug("false")
			}
		}
	}
}

func (m *manager) generateSessionID() string {
	return uuid.New().String()
}

func (m *manager) doesDelete() bool {
	if m.probOfDelete <= 0 {
		return false
	}
	if m.probOfDeleteDivisor <= m.probOfDelete {
		return true
	}

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(m.probOfDeleteDivisor) + 1

	log.Debug(strconv.Itoa(n))

	return n <= m.probOfDelete
}
