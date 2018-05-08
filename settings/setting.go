package settings

import "time"

func init() {
	Session = &session{
		Name:      "GOSESSID",
		ExpiresIn: 1 * time.Hour,
	}
}
