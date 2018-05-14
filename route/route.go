package route

import (
	"net/http"

	"github.com/tomocy/goron/handlers/count"
)

type Route interface {
	Listen()
}

type route struct {
}

func New() Route {
	return &route{}
}

func (r *route) Listen() {
	// Write root
	http.HandleFunc("/count", count.Index)
}
