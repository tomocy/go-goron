package count

import (
	"net/http"

	"github.com/tomocy/goron/session/manager"
	"github.com/tomocy/goron/template"
)

var sessionManager manager.Manager
var tmpl template.Template

func Index(w http.ResponseWriter, r *http.Request) {
	session := sessionManager.GetSession(w, r)

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

func init() {
	tmpl = template.New()
	sessionManager, _ = manager.New("memory")
}
