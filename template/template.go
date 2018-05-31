package template

import (
	tmpl "html/template"
	"io"
)

var templates map[string]*tmpl.Template

type Template interface {
	Render(w io.Writer, name string, data interface{}) error
}

type template struct {
}

func init() {
	base := "views/layouts/master.html"
	templates = make(map[string]*tmpl.Template)
	templates["error"] = tmpl.Must(tmpl.ParseFiles(base, "views/layouts/error.html"))
	templates["countIndex"] = tmpl.Must(tmpl.ParseFiles(base, "views/count.html"))
	templates["testIndex"] = tmpl.Must(tmpl.ParseFiles(base, "views/test.html"))
}

func New() Template {
	return &template{}
}

func (t *template) Render(w io.Writer, name string, data interface{}) error {
	return templates[name].ExecuteTemplate(w, "master.html", data)
}
