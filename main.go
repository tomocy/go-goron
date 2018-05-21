package main

import (
	"os"

	"github.com/tomocy/goron/route"
	"github.com/tomocy/goron/session/manager"
)

func main() {
	rtCh := make(chan os.Signal, 1)
	r := route.New(rtCh)
	m, err := manager.New("file")
	if err != nil {
		panic(err)
	}

	go r.ListenAndServe()
	go m.DeleteExpiredSessions()

	<-rtCh
}
