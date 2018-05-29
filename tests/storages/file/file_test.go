package file

import (
	"reflect"
	"testing"
	"time"

	"github.com/tomocy/goron/session/storages/file"
)

const (
	Delimiter    = ":"
	ExpiresAtKey = "expiresAt"
	TimeLayout   = time.RFC3339Nano
)

func TestInitSession(t *testing.T) {
	f := file.New()
	sess1ID := generateSessionID()

	// functino to be tested
	sess1 := f.InitSession(sess1ID)
	if sess1ID != sess1.ID() {
		t.Errorf("could not init session with passed id\nwanted: %s\nhad: %s", sess1ID, sess1.ID())
	}

	sess2, err := f.GetSession(sess1ID)
	if err != nil {
		t.Fatal("could not get session", err)
	}

	if sess1.ID() != sess2.ID() {
		t.Errorf("session id not same\nwanted: %s\nhad: %s", sess1.ID(), sess2.ID())
	}
	if !reflect.DeepEqual(sess1.Data(), sess2.Data()) {
		t.Errorf("data in session not same\nwanted: %#v\nhad: %#v", sess1.Data(), sess2.Data())
	}
	if !sess1.ExpiresAt().Equal(sess2.ExpiresAt()) {
		t.Error("expires of session not same")
	}
}
