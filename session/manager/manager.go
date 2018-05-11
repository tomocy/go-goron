package manager

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tomocy/goron/session"
	"github.com/tomocy/goron/session/storages"
)

type Manager interface {
	CreateSession() session.Session
	GetSession(string) (session.Session, error)
	SetSession(string, map[string]interface{})
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

func (m *manager) CreateSession() session.Session {
	sessionID := uuid.New().String()

	return m.storage.InitSession(sessionID)
}

func (m *manager) GetSession(sessionID string) (session.Session, error) {
	session, err := m.storage.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (m *manager) SetSession(sessionID string, dat map[string]interface{}) {
	m.storage.SetSession(sessionID, dat)
}
