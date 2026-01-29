package api

import (
	"net/http"
)

func (s *Server) registerRoutes() {
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/ws", s.handleWebsocket)
	http.HandleFunc("/health", s.handleHealthCheck)
}
