package settings

import "time"

type cookie struct {
	ExpiresIn time.Duration
}

var Cookie *cookie

func (c *cookie) SetExpiresIn(expIn time.Duration) {
	c.ExpiresIn = expIn
}

func init() {
	Cookie = &cookie{
		ExpiresIn: 1 * time.Hour,
	}
}
