package file

import "github.com/google/uuid"

func generateSessionID() string {
	return uuid.New().String()
}
