package main

import (
	"github.com/tomocy/goron/route"
)

var r route.Route

func main() {
	r.ListenAndServe()
}

func init() {
	r = route.New()
}
