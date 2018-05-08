package settings

import "time"

type session struct {
	Name      string
	ExpiresIn time.Duration
}

var Session *session
