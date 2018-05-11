package main

import (
	"fmt"
	"net/http"

	"github.com/tomocy/goron/session/cookie"
	"github.com/tomocy/goron/session/manager"
	"github.com/tomocy/goron/template"
)

var sessionManager manager.Manager
var tmpl template.Template

func init() {
	tmpl = template.New()
	sessionManager, _ = manager.New("memory")
}

func main() {
	http.HandleFunc("/count", countIndex)

	fmt.Println("Listeing :8080 ...")
	http.ListenAndServe(":8080", nil)
}

func countIndex(w http.ResponseWriter, r *http.Request) {
	sessionID, err := cookie.GetSessionID(r)
	if err != nil {
		session := sessionManager.CreateSession()
		sessionID = session.ID()

		cookie.SetSessionID(w, sessionID)
	}

	session, err := sessionManager.GetSession(sessionID)
	if err != nil {
		// In this case, the storage should be memory
		// and cookie still remains while session in the memory is empty
		// So delete the session id in cookie, and recreate session and reset sessino id in cookie
		cookie.DestroySessionID(w)
		session = sessionManager.CreateSession()
		sessionID = session.ID()

		cookie.SetSessionID(w, sessionID)
	}

	cnt, ok := session.Get("count").(int)
	if !ok {
		cnt = 1
	} else {
		cnt++
	}

	session.Set("count", cnt)

	dat := struct {
		Cnt interface{}
	}{
		Cnt: cnt,
	}

	tmpl.Render(w, "countIndex", dat)
}
