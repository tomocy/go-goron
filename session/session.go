package session

import (
	"sync"
	"time"

	"github.com/tomocy/goron/log"
	"github.com/tomocy/goron/settings"
)

type Session interface {
	Set(k string, v string)
	Get(k string) string
	ID() string
	ExpiresAt() time.Time
	Data() map[string]string
	DoesExpire() bool
}

type session struct {
	id        string
	dat       map[string]string
	expiresAt time.Time
	mu        sync.Mutex
}

func New(id string, dat map[string]string) Session {
	expiresAt := time.Now().Add(settings.Session.ExpiresIn)

	return &session{id: id, dat: dat, expiresAt: expiresAt}
}

func (s *session) Set(k string, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.dat[k] = v
}

func (s *session) Get(k string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.dat[k]
	if !ok {
		return ""
	}

	return v
}

func (s *session) ID() string {
	return s.id
}

func (s *session) ExpiresAt() time.Time {
	return s.expiresAt
}

func (s *session) Data() map[string]string {
	return s.dat
}

func (s *session) DoesExpire() bool {
	log.Debug(s.expiresAt.String())
	if s.expiresAt.Before(time.Now()) {
		return true
	}

	return false
}
