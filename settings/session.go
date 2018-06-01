package settings

import "time"

type session struct {
	Name      string
	ExpiresIn time.Duration
	Storage   string
}

var Session *session

func (s *session) SetName(name string) {
	s.Name = name
}

func (s *session) SetExpiresIn(expIn time.Duration) {
	s.ExpiresIn = expIn
}

func (s *session) SetStorage(strg string) {
	s.Storage = strg
}

func init() {
	Session = &session{
		Name:      "GOSESSID",
		ExpiresIn: 30 * time.Minute,
		Storage:   "file",
	}
}
