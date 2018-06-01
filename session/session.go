package session

import (
	"sync"
	"time"
)

// type Session interface {
// 	Set(k string, v string)
// 	Get(k string) string
// 	ID() string
// 	ExpiresAt() time.Time
// 	Data() map[string]string
// 	DoesExpire() bool
// }

type Session struct {
	id        string
	dat       map[string]string
	expiresAt time.Time
	mu        sync.Mutex
}

func New(id string, expiresAt time.Time, dat map[string]string) *Session {
	// create new map and copy dat to it
	newDat := make(map[string]string)
	for k, v := range dat {
		newDat[k] = v
	}
	return &Session{id: id, dat: newDat, expiresAt: expiresAt}
}

func (s *Session) Set(k string, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.dat[k] = v
}

func (s *Session) Get(k string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.dat[k]
	if !ok {
		return ""
	}

	return v
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) ExpiresAt() time.Time {
	return s.expiresAt
}

func (s *Session) Data() map[string]string {
	return s.dat
}

func (s *Session) DoesExpire() bool {
	if s.expiresAt.Before(time.Now()) {
		return true
	}

	return false
}
