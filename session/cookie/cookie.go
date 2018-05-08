package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/tomocy/goron/settings"
)

func SetSessionID(c echo.Context, sessionID string) {
	cookie := &http.Cookie{
		Name:    settings.Session.Name,
		Value:   sessionID,
		Expires: time.Now().Add(settings.Session.ExpiresIn),
	}

	c.SetCookie(cookie)
}
func GetSessionID(c echo.Context) (string, error) {
	cookie, err := c.Cookie(settings.Session.Name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
