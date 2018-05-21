package storages

import (
	"errors"

	"github.com/tomocy/goron/session"
	"github.com/tomocy/goron/session/storages/file"
	"github.com/tomocy/goron/session/storages/memory"
)

type Storage interface {
	InitSession(sessionID string) session.Session
	GetSession(sessionID string) (session.Session, error)
	SetSession(session session.Session)
	DeleteSession(sessionID string)
	DeleteExpiredSessions()
}

func Get(storage string) (Storage, error) {
	switch storage {
	case "memory":
		return memory.New(), nil
	case "file":
		return file.New(), nil
	default:
		return nil, errors.New("Not found")
	}
}
