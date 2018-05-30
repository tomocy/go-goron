package manager

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/tomocy/goron/session/manager"
	"github.com/tomocy/goron/session/storages/file"
	"github.com/tomocy/goron/settings"
)

type managerTest struct {
	m   manager.Manager
	rec *httptest.ResponseRecorder
	req *http.Request
}

func TestMain(m *testing.M) {
	ec := m.Run()
	tearDown()
	os.Exit(ec)
}

func tearDown() {
	os.RemoveAll("storage")
}

func TestGetSession(t *testing.T) {
	t.Run("No cookie", testOnNoCookie)
	t.Run("No session while cookie exists", testOnNoSession)
	t.Run("Session expires", testOnSessionExpired)
}

func testOnNoCookie(t *testing.T) {
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

func testOnNoSession(t *testing.T) {
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

func testOnSessionExpired(t *testing.T) {
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
func TestSetSession(t *testing.T) {
	strg := file.New()
	sess1ID := "test"
	sess1 := strg.InitSession(sess1ID)
	dat1 := map[string]string{
		"aiueo": "test",
		"あいうえお": "アイウエオ",
	}

	for k, v := range dat1 {
		sess1.Set(k, v)
	}

	m, err := manager.New(settings.Session.Storage)
	if err != nil {
		t.Fatal(err)
	}
	m.SetSession(sess1)

	sess2, err := strg.GetSession(sess1ID)
	if err != nil {
		t.Fatal(err)
	}
	sess2ID := sess2.ID()
	dat2 := sess2.Data()

	if sess1ID != sess2ID {
		t.Errorf("session id\n\twanted %s\n\thad %s", sess1ID, sess2ID)
	}
	if !reflect.DeepEqual(dat1, dat2) {
		t.Errorf("sessino data\n\twanted %#v\n\thad %#v", dat1, dat2)
	}
	if !sess1.ExpiresAt().Equal(sess2.ExpiresAt()) {
		t.Errorf("expires times of session1 and session2 are not same")
	}
}

func TestDeleteExpiredSessions(t *testing.T) {
	m, err := manager.New(settings.Session.Storage)
	if err != nil {
		t.Fatal(err)
	}

	go m.DeleteExpiredSessions()

	time.Sleep(1 * time.Second)
}
