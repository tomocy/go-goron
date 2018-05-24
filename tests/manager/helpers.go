package manager

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/tomocy/goron/session"
	"github.com/tomocy/goron/session/manager"
	"github.com/tomocy/goron/session/storages/file"
	"github.com/tomocy/goron/settings"
)

func onNoCookie(t *testing.T) {
	mt, err := setUpManagerTest()
	if err != nil {
		t.Fatal(err)
	}

	sess1 := mt.m.GetSession(mt.rec, mt.req)
	sess2 := mt.m.GetSession(mt.rec, mt.req)

	if sess1.ID() == sess2.ID() {
		t.Error("Session ids are same mainly because manager cannot recreate session when no session id in cookie")
	}
}

func onNoSession(t *testing.T) {
	mt, err := setUpManagerTest()
	if err != nil {
		t.Fatal(err)
	}

	strg := file.New()
	sess1 := strg.InitSession("test")

	mt.req.AddCookie(&http.Cookie{
		Name:    settings.Session.Name,
		Value:   sess1.ID(),
		Expires: sess1.ExpiresAt(),
	})

	err = deleteSession(sess1.ID())
	if err != nil {
		t.Fatal(err)
	}

	sess2 := mt.m.GetSession(mt.rec, mt.req)

	if sess1.ID() == sess2.ID() {
		t.Errorf("Session ids are same mainly because manager could not recreate session though no the session in server")
	}
}

func onSessionExpired(t *testing.T) {
	mt, err := setUpManagerTest()
	if err != nil {
		t.Fatal(err)
	}

	strg := file.New()
	sess1 := strg.InitSession("test")

	makeSessionExpires(mt.m, sess1)

	sess2 := mt.m.GetSession(mt.rec, mt.req)

	if sess1.ID() == sess2.ID() {
		t.Error("Session ids are same mainly because manager could not recreate session though the former session expires")
	}
}

func setUpManagerTest() (*managerTest, error) {
	m, err := manager.New(settings.Session.Storage)
	if err != nil {
		return nil, err
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://192.168.55.55:8080/count", nil)
	if err != nil {
		return nil, err
	}

	return &managerTest{
		m:   m,
		rec: rec,
		req: req,
	}, nil
}

func deleteSession(sessID string) error {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	fname := "storage/sessions/" + sessID
	err := os.Remove(fname)

	return err
}

func makeSessionExpires(m manager.Manager, sess session.Session) {
	expiredSess := session.New(sess.ID(), time.Now().Add(-1*time.Hour), sess.Data())
	m.SetSession(expiredSess)
}
