package api

import (
	"log"
	"fmt"
	"net/http"
)

type Server struct {
	port string
}

func NewServer(port string) Server {
	log.Printf("Creating server\n")
	return Server{port: port}
}

func (s *Server) Run() error {
	log.Printf("Server is listening on port %s...\n", s.port)
	s.registerRoutes()
	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
}
