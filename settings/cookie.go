package settings

import "time"

type cookie struct {
	expiresIn time.Duration
}

var Cookie *cookie

func (c *cookie) ExpiresIn() time.Duration {
	return c.expiresIn
}

func (c *cookie) SetExpiresIn(expIn time.Duration) {
	c.expiresIn = expIn
}

func init() {
	Cookie = &cookie{
		expiresIn: 1 * time.Hour,
	}
}
