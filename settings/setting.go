package settings

import "time"

func init() {
	Session = &session{
		Name:      "GOSESSID",
		ExpiresIn: 30 * time.Minute,
		Storage:   "file",
	}

	SessionManager = &sessionManager{
		ProbOfDelete:        1,
		ProbOfDeleteDivisor: 100,
	}

	Cookie = &cookie{
		ExpiresIn: 1 * time.Hour,
	}
}
