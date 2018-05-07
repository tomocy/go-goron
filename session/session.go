package session

import "time"

type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	ID() string
}

type session struct {
	id        Id
	dat       map[string]interface{}
	expiresAt time.Time
}

type Id string

const expiresIn time.Duration = (3 * time.Minute)

func (s *session) Set(key string, value interface{}) error {
	return nil
}

func (s *session) Get(key string) (interface{}, error) {
	return nil, nil
}

func (s *session) Delete(key string) error {
	return nil
}

func (s *session) ID() Id {
	return s.id
}
