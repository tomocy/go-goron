package file

import (
	"reflect"
	"testing"
	"time"

	"github.com/tomocy/goron/log/tlog"
	"github.com/tomocy/goron/session/storages/file"
	"github.com/tomocy/goron/settings"
)

const (
	Delimiter    = ":"
	ExpiresAtKey = "expiresAt"
	TimeLayout   = time.RFC3339Nano
)

func TestInitSession(t *testing.T) {
	f := file.New()
	sessID := generateSessionID()

	// functino to be tested
	sess := f.InitSession(sessID)

	if sessID != sess.ID() {
		t.Error(tlog.GetWantedHad("could not init session with passed id", sessID, sess.ID()))
	}

	emptData := make(map[string]string)
	if !reflect.DeepEqual(emptData, sess.Data()) {
		t.Error(tlog.GetWantedHad("could not init session with empty data", emptData, sess.Data()))
	}

	if sess.ExpiresAt().Add(-1 * settings.Session.ExpiresIn).After(time.Now()) {
		t.Error("could not set session expires as settings")
	}
}
