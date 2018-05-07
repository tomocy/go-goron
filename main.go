package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var tmpls map[string]*template.Template

type Template struct {
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return tmpls[name].ExecuteTemplate(w, "master.html", data)
}

func init() {
	loadTemplates()
}
func main() {
	e := echo.New()

	t := &Template{}
	e.Renderer = t

	e.Use(middleware.Logger(), middleware.Recover())

	e.GET("/", index)

	e.Logger.Fatal(e.Start(":8080"))
}

func loadTemplates() {
	var baseTmpl = "views/layouts/master.html"
	tmpls = make(map[string]*template.Template)
	tmpls["index"] = template.Must(template.ParseFiles(baseTmpl, "views/hello.html"))
}

func index(c echo.Context) error {
	dat := struct {
		World string
	}{
		World: "aiueo",
	}

	return c.Render(http.StatusOK, "index", dat)
}
