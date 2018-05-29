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

func TestGetSession(t *testing.T) {
	f := file.New()
	sess1ID := generateSessionID()
	sess1 := f.InitSession(sess1ID)

	// function to be tested
	sess2, err := f.GetSession(sess1ID)
	if err != nil {
		t.Fatal("could not get session\n", err)
	}

	if sess1.ID() != sess2.ID() {
		t.Error(tlog.GetWantedHad("session id not same", sess1.ID(), sess2.ID()))
	}
	if !reflect.DeepEqual(sess1.Data(), sess2.Data()) {
		t.Error(tlog.GetWantedHad("data in session not same", sess1.Data(), sess2.Data()))
	}
	if !sess1.ExpiresAt().Equal(sess2.ExpiresAt()) {
		t.Error("expires of session not same")
	}

}

func TestSetSession(t *testing.T) {
	f := file.New()
	sess1ID := generateSessionID()
	sess1 := f.InitSession(sess1ID)

	dat := map[string]string{
		"aiueo":       "あいうえお",
		"kakikukeko":  "かきくけこ",
		"sashisuseso": "さしすせそ",
		"tatituteto":  "たちつてと",
		"hahihuheho":  "はひふへほ",
		"mamimumemo":  "まみむめも",
		"yayuyo":      "やゆよ",
		"rarirurero":  "らりるれろ",
		"waronn":      "わをん",
	}
	for k, v := range dat {
		sess1.Set(k, v)
	}

	// function to be tested
	f.SetSession(sess1)

	sess2, err := f.GetSession(sess1ID)
	if err != nil {
		t.Fatal("could not get session")
	}

	if sess1.ID() != sess2.ID() || !reflect.DeepEqual(dat, sess2.Data()) {
		t.Error(tlog.GetWantedHad("data in session not same", dat, sess2.Data()))
	}
}
