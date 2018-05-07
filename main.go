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

	e.GET("/greet/create", greetNew)
	e.POST("/greet/create", greetCreate)

	e.Logger.Fatal(e.Start(":8080"))
}

func loadTemplates() {
	var baseTmpl = "views/layouts/master.html"
	tmpls = make(map[string]*template.Template)
	tmpls["greetIndex"] = template.Must(template.ParseFiles(baseTmpl, "views/greet/index.html"))
	tmpls["greetCreate"] = template.Must(template.ParseFiles(baseTmpl, "views/greet/create.html"))
}

func greetNew(c echo.Context) error {
	return c.Render(http.StatusOK, "greetCreate", nil)
}

func greetCreate(c echo.Context) error {
	obj := c.FormValue("to")

	return c.Render(http.StatusOK, "greetIndex", obj)
}
