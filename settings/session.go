package settings

import "time"

type session struct {
	name      string
	expiresIn time.Duration
	storage   string
}

var Session *session

func (s *session) Name() string {
	return s.name
}

func (s *session) ExpiresIn() time.Duration {
	return s.expiresIn
}

func (s *session) Storage() string {
	return s.storage
}

func (s *session) SetName(name string) {
	s.name = name
}

func (s *session) SetExpiresIn(expIn time.Duration) {
	s.expiresIn = expIn
}

func (s *session) SetStorage(strg string) {
	s.storage = strg
}

func init() {
	Session = &session{
		name:      "GOSESSID",
		expiresIn: 30 * time.Minute,
		storage:   "file",
	}
}
