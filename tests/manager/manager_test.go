package manager

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/tomocy/goron/session"
	"github.com/tomocy/goron/session/manager"
	"github.com/tomocy/goron/settings"
)

func TestMain(m *testing.M) {
	ec := m.Run()
	os.Exit(ec)
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
		t.Error("wanted type of s was Manager, but had %s", tp)
	}
}
