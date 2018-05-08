package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

var tmpls map[string]*template.Template

type Template struct {
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return tmpls[name].ExecuteTemplate(w, "master.html", data)
}

func initTemplates() {
	var baseTmpl = "views/layouts/master.html"
	tmpls = make(map[string]*template.Template)
	// tmpls["greetIndex"] = template.Must(template.ParseFiles(baseTmpl, "views/greet/index.html"))
	// tmpls["greetCreate"] = template.Must(template.ParseFiles(baseTmpl, "views/greet/create.html"))
	tmpls["error"] = template.Must(template.ParseFiles(baseTmpl, "views/layouts/error.html"))
	tmpls["countIndex"] = template.Must(template.ParseFiles(baseTmpl, "views/count.html"))
}

// func greetNew(c echo.Context) error {
// 	return c.Render(http.StatusOK, "greetCreate", nil)
// }

// func greetCreate(c echo.Context) error {
// 	obj := c.FormValue("to")
// 	sessionId, err := sessionManager.CreateSession()
// 	if err != nil {
// 		return c.Render(http.StatusInternalServerError, "error", err.Error())
// 	}

// 	dat := map[string]interface{}{"to": obj}
// 	if err := sessionManager.SaveSession(sessionId, dat); err != nil {
// 		return c.Render(http.StatusInternalServerError, "error", err.Error())
// 	}

// 	log.Fatal(sessionId)

// 	return c.Render(http.StatusOK, "greetIndex", obj)
// }
