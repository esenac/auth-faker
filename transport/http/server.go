package http

import (
	"fmt"
	"net/http"
)

// Server is the HTTP server.
type Server struct {
	mux *http.ServeMux
}

// New creates a Server.
func New() Server {
	mux := http.NewServeMux()
	return Server{mux}
}

// Start runs the server.
func (s Server) Start(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.mux)
}

// AddRoute adds a route to the Server, managed by the provided handler.
func (s Server) AddRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(route, handler)
}
