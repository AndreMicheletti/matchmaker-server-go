package api

import (
	"fmt"
	"net/http"
	"html/template"
)

func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
	tmpl, err := template.ParseFiles("internal/api/templates/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}
