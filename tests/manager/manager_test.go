package manager

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/tomocy/goron/session"
	"github.com/tomocy/goron/session/manager"
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
	t.Run("No cookie", onNoCookie)
	t.Run("No session", onNoSession)
	t.Run("Session expires", onSessionExpired)
}

func TestDeleteExpiredSessions(t *testing.T) {
	m, err := manager.New(settings.Session.Storage)
	if err != nil {
		t.Fatal(err)
	}

	go m.DeleteExpiredSessions()

	time.Sleep(1 * time.Second)
}
