package http

import (
	"fmt"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func New() Server {
	mux := http.NewServeMux()
	return Server{mux}
}

func (s Server) Start(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.mux)
}

func (s Server) AddRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(route, handler)
}
