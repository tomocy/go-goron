package main

import (
	"os"

	"github.com/tomocy/goron/route"
	"github.com/tomocy/goron/session/manager"
	"github.com/tomocy/goron/settings"
)

func main() {
	rtCh := make(chan os.Signal, 1)
	r := route.New(rtCh)
	m, err := manager.New(settings.Session.Storage)
	if err != nil {
		panic(err)
	}

	go r.ListenAndServe()
	go m.DeleteExpiredSessions()

	<-rtCh
}
