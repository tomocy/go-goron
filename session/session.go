package session

import "sync"

type Session interface {
	Set(k string, v string)
	Get(k string) string
	ID() string
	Data() map[string]string
}

type session struct {
	id  string
	dat map[string]string
	mu  sync.Mutex
}

func New(id string, dat map[string]string) Session {
	return &session{id: id, dat: dat}
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

func (s *session) Data() map[string]string {
	return s.dat
}
