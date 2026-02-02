package api

import (
	"log"
	"fmt"
	"net/http"
	"github.com/AndreMicheletti/matchmaker-server-go/internal/matchmaker"
)

type Server struct {
	port string
	eng *matchmaker.Engine
}

func NewServer(port string, engine matchmaker.Engine) Server {
	log.Printf("Creating server\n")
	return Server{port: port, eng: &engine}
}

func (s *Server) Run() error {
	log.Printf("Server is listening on port %s...\n", s.port)
	s.registerRoutes()
	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
}

func (s *Server) Engine() *matchmaker.Engine {
	return s.eng
}

func (s *Server) Close() {
	s.eng.Close()
}
