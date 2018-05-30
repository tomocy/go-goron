package helper

import (
	"github.com/google/uuid"
)

func GenerateSessionID() string {
	return uuid.New().String()
}
