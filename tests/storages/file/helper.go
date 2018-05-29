package file

import "github.com/google/uuid"

func generateSessionID() string {
	return uuid.New().String()
}

func getSessionFilePath(sessionID string) string {
	b := make([]byte, 0, 10)
	b = append(b, "storage/sessions/"...)
	b = append(b, sessionID...)

	return string(b)
}
