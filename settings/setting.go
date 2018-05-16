package settings

import "time"

func init() {
	Session = &session{
		Name:      "GOSESSID",
		ExpiresIn: 30 * time.Minute,
		Storage:   "file",
	}

	Cookie = &cookie{
		ExpiresIn: 1 * time.Hour,
	}
}
