package main

import (
	"os"

	"github.com/tomocy/goron/route"
)

func main() {
	rtCh := make(chan os.Signal, 1)
	r := route.New(rtCh)
	go r.ListenAndServe()
	<-rtCh
}
