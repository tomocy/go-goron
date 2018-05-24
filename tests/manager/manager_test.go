package manager

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/tomocy/goron/session"
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
	m, err := manager.New(settings.Session.Storage)
	if err != nil {
		t.Fatalf("faild to get new manager: %s", err)
	}

	// request and response
	req := httptest.NewRequest("GET", "http://192.168.55.55:8080/count", nil)
	res := httptest.NewRecorder()

	s := m.GetSession(res, req)

	switch tp := s.(type) {
	case session.Session:
	default:
		t.Errorf("wanted type of s was Manager, but had %s", tp)
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
