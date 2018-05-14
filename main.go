package main

import (
	"fmt"
	"net/http"

	"github.com/tomocy/goron/route"
)

var r route.Route

func main() {
	r.Listen()

	fmt.Println("Listeing :8080")
	http.ListenAndServe(":8080", nil)
}

func init() {
	r = route.New()
}
