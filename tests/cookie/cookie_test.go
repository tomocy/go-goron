package cookie

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/tomocy/goron/helper"
	"github.com/tomocy/goron/log/tlog"

	"github.com/tomocy/goron/session/cookie"
	"github.com/tomocy/goron/settings"
)

func TestGetSessionID(t *testing.T) {
	req := httptest.NewRequest("GET", "http://192.168.55.55:8080/count", nil)
	sessID1 := helper.GenerateSessionID()
	req.AddCookie(&http.Cookie{
		Name:    settings.Session.Name(),
		Value:   sessID1,
		Expires: time.Now().Add(settings.Session.ExpiresIn()),
	})

	// function to be tested
	sessID2, err := cookie.GetSessionID(req)
	if err != nil {
		t.Fatal("could not get session id from cookie\n", err)
	}

	if sessID1 != sessID2 {
		t.Errorf(tlog.GetWantedHad("session id not expected", sessID1, sessID2))
	}
}

func TestSetSessionID(t *testing.T) {
	sessID := helper.GenerateSessionID()
	rec := httptest.NewRecorder()

	// function to be tested
	cookie.SetSessionID(rec, sessID)

	setCookie := rec.HeaderMap.Get("Set-Cookie")
	sessInfo := getSessionInfo(sessID)
	if !strings.Contains(setCookie, sessInfo) {
		t.Errorf(tlog.GetWantedHad("session info not expected", sessInfo, setCookie))
	}
}
