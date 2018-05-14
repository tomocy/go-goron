package settings

import "time"

type session struct {
	Name      string
	ExpiresIn time.Duration
	Storage   string
}

var Session *session
