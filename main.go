package main

import (
	"log"
	"net/http"

	"github.com/tomocy/goron/session"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var s session.Session

func init() {
	initTemplates()
	s = session.New()
}

func main() {
	e := echo.New()

	t := &Template{}
	e.Renderer = t

	e.Use(middleware.Logger(), middleware.Recover())

	// e.GET("/greet/create", greetNew)
	// e.POST("/greet/create", greetCreate)

	e.GET("/count", countIndex)

	e.Start(":8080")
}

func countIndex(c echo.Context) error {
	cnt := s.Get("count")
	if cnt == nil {
		s.Set("count", 0)
		cnt = 0
	} else {
		s.Set("count", cnt.(int)+1)

	}

	log.Println(cnt)

	dat := struct {
		Cnt interface{}
	}{
		Cnt: s.Get("count"),
	}
	return c.Render(http.StatusOK, "countIndex", dat)
}
