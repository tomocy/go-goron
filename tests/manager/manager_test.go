package manager

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sync"
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

func onNoCookie(t *testing.T) {
	mt, err := setUpManagerTest()
	if err != nil {
		t.Fatal(err)
	}

	sess1 := mt.m.GetSession(mt.rec, mt.req)

	sess1.Set("count", "0")
	mt.m.SetSession(sess1)

	sess2 := mt.m.GetSession(mt.rec, mt.req)

	if !reflect.DeepEqual(sess1.Data(), sess2.Data()) {
		t.Errorf("\t\nwanted %#v,\t\nhad %#v", sess1.Data(), sess2.Data())
	}
}

func onNoSession(t *testing.T) {
	mt, err := setUpManagerTest()
	if err != nil {
		t.Fatal(err)
	}

	sess := mt.m.GetSession(mt.rec, mt.req)

	mt.req.AddCookie(&http.Cookie{
		Name:    settings.Session.Name,
		Value:   sess.ID(),
		Expires: sess.ExpiresAt(),
	})

	err = deleteSession(sess.ID())
	if err != nil {
		t.Fatal(err)
	}

	sess1 := mt.m.GetSession(mt.rec, mt.req)
	mt.req.Header.Del("Cookie")
	mt.req.AddCookie(&http.Cookie{
		Name:    settings.Session.Name,
		Value:   sess1.ID(),
		Expires: sess1.ExpiresAt(),
	})

	sess1.Set("count", "0")
	mt.m.SetSession(sess1)

	sess2 := mt.m.GetSession(mt.rec, mt.req)

	if !reflect.DeepEqual(sess1.Data(), sess2.Data()) {
		t.Errorf("\t\nwanted %#v,\t\nhad %#v", sess1.Data(), sess2.Data())
	}
}

func onSessionExpired(t *testing.T) {
	mt, err := setUpManagerTest()
	if err != nil {
		t.Fatal(err)
	}

	sess1 := mt.m.GetSession(mt.rec, mt.req)

	sess1.Set("count", "0")
	mt.m.SetSession(sess1)

	makeSessionExpires(mt.m, sess1)

	sess2 := mt.m.GetSession(mt.rec, mt.req)

	if sess1.ID() == sess2.ID() {
		t.Error("sess1's id and sess2's id are same, this may be because manager could not check if session expires conrectly")
	}
}

func setUpManagerTest() (*managerTest, error) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	}))
	defer s.Close()

	m, err := manager.New(settings.Session.Storage)
	if err != nil {
		return nil, err
	}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", s.URL+"/count", nil)
	if err != nil {
		return nil, err
	}

	sess := m.GetSession(rec, req)

	req.AddCookie(&http.Cookie{
		Name:    settings.Session.Name,
		Value:   sess.ID(),
		Expires: sess.ExpiresAt(),
	})

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
