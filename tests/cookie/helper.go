package cookie

import (
	"github.com/tomocy/goron/settings"
)

func getSessionInfo(sessionID string) string {
	b := make([]byte, 0, 10)
	b = append(b, settings.Session.Name...)
	b = append(b, "="...)
	b = append(b, sessionID...)

	return string(b)
}
