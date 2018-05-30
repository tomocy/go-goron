package session

import (
	"sync"
	"time"
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

func New(id string, expiresAt time.Time, dat map[string]string) Session {
	// allocate new memroy to data
	newDat := make(map[string]string)
	for k, v := range dat {
		newDat[k] = v
	}
	return &session{id: id, dat: newDat, expiresAt: expiresAt}
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
	if s.expiresAt.Before(time.Now()) {
		return true
	}

	return false
}
