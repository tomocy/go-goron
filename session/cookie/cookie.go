package cookie

import (
	"net/http"
	"time"

	"github.com/tomocy/goron/settings"
)

func SetSessionID(w http.ResponseWriter, sessionID string) {
	cookie := &http.Cookie{
		Name:    settings.Session.Name,
		Value:   sessionID,
		Expires: time.Now().Add(settings.Cookie.ExpiresIn),
	}

	http.SetCookie(w, cookie)
}

func GetSessionID(r *http.Request) (string, error) {
	cookie, err := r.Cookie(settings.Session.Name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
