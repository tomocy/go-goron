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

	serve()
}

type route struct {
	stopCh chan os.Signal
}

func New() Route {
	stopCh := make(chan os.Signal, 1)
	return &route{stopCh: stopCh}
}

func (r *route) ListenAndServe() {
	// Write root
	http.HandleFunc("/count", count.Index)

	// And serve
	r.serve()
}

func (r *route) serve() {
	fmt.Println("Listeing :8080")
	http.ListenAndServe(":8080", nil)
	signal.Notify(r.stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-r.stopCh
}
