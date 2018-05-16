package settings

import "time"

type cookie struct {
	ExpiresIn time.Duration
}

var Cookie *cookie
