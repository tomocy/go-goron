package cookie

import (
	"github.com/google/uuid"
	"github.com/tomocy/goron/settings"
)

func generateSessionID() string {
	return uuid.New().String()
}

func getSessionInfo(sessionID string) string {
	b := make([]byte, 0, 10)
	b = append(b, settings.Session.Name...)
	b = append(b, "="...)
	b = append(b, sessionID...)

	return string(b)
}
