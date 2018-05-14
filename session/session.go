package session

import "sync"

type Session interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	ID() string
}

type session struct {
	id  string
	dat map[string]interface{}
	mu  sync.Mutex
}

func New(id string, dat map[string]interface{}) Session {
	return &session{id: id, dat: dat}
}

func (s *session) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.dat[key] = value
}

func (s *session) Get(key string) interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.dat[key]
	if !ok {
		return nil
	}

	return v
}

func (s *session) ID() string {
	return s.id
}
