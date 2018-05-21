package count

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tomocy/goron/session/manager"
	"github.com/tomocy/goron/settings"
	"github.com/tomocy/goron/template"
)

var sessionManager manager.Manager
var tmpl template.Template

func Index(w http.ResponseWriter, r *http.Request) {
	session := sessionManager.GetSession(w, r)

	cnt, err := strconv.Atoi(session.Get("count"))
	if err != nil {
		cnt = 1
	} else {
		cnt++
	}

	session.Set("count", fmt.Sprintf("%d", cnt))
	sessionManager.SetSession(session)

	dat := struct {
		Cnt interface{}
	}{
		Cnt: cnt,
	}

	tmpl.Render(w, "countIndex", dat)
}

func init() {
	tmpl = template.New()

	var err error
	sessionManager, err = manager.New(settings.Session.Storage)
	if err != nil {
		panic(err)
	}
}
