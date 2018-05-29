package session

import (
	"reflect"
	"testing"
	"time"

	"github.com/tomocy/goron/log/tlog"

	"github.com/tomocy/goron/session"
)

func TestNew(t *testing.T) {
	sessID := generateSessionID()
	expiresAt := time.Now()
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

	// function to be tested
	sess := session.New(sessID, expiresAt, dat)

	if sessID != sess.ID() {
		t.Error(tlog.GetWantedHad("id not same", sessID, sess.ID()))
	}
	if !reflect.DeepEqual(dat, sess.Data()) {
		t.Error(tlog.GetWantedHad("data in session not same", dat, sess.Data()))
	}
	if !expiresAt.Equal(sess.ExpiresAt()) {
		t.Error(tlog.GetWantedHad("expiresAt of session not same", expiresAt, sess.ExpiresAt()))
	}
}

func TestGet(t *testing.T) {
	sessID := generateSessionID()
	expiresAt := time.Now()
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
	sess := session.New(sessID, expiresAt, dat)

	for k, v1 := range dat {
		// function to be tested
		v2 := sess.Get(k)
		if v1 != v2 {
			t.Error(tlog.GetWantedHad("value of "+k+" not expected", v1, v2))
		}
	}
}
