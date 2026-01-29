package api

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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

func (s *Server) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("[WEBSOCKET] connection received")
	outCh := make(chan []byte, 256)
	defer conn.Close()
	defer close(outCh)
	go writeMessage(conn, outCh, websocket.TextMessage)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[WEBSOCKET] message received: %s\n", p)
		outCh <- p
	}
}

func writeMessage(conn *websocket.Conn, outChannel chan []byte, msgType int) {
	for msg := range outChannel {
		log.Printf("[WEBSOCKET] returning message: %s\n", msg)
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
		}
	}
}