package route

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/tomocy/goron/handlers/count"
)

type Route interface {
	ListenAndServe()

	listen()
	serve()
}

type route struct {
	stopCh chan os.Signal
}

func New(stopCh chan os.Signal) Route {
	// stopCh := make(chan os.Signal, 1)
	return &route{stopCh: stopCh}
}

func (r *route) ListenAndServe() {
	// Listen
	r.listen()

	// And serve
	r.serve()
}

func (r *route) listen() {
	// Routing
	http.HandleFunc("/count", count.Index)
}

func (r *route) serve() {
	fmt.Println("Listeing :8080")
	http.ListenAndServe(":8080", nil)
	signal.Notify(r.stopCh, syscall.SIGINT, syscall.SIGTERM)
}
