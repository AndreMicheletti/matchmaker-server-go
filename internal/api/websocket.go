package api

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID string
	Conn *websocket.Conn
	SendCh chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *Client) writePump() {
	for msg := range c.SendCh {
		log.Printf("[WEBSOCKET] returning message: %s\n", msg)
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println(err)
		}
	}
}

func (c *Client) readPump() {
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[WEBSOCKET] message received: %s\n", msg)
		c.SendCh <- msg
	}
}

func (s *Server) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID := query.Get("user_id")

	if userID == "" {
		log.Println("[WEBSOCKET] user_id não fornecido")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("[WEBSOCKET] connection received for user %s", userID)

	c := Client{}
	c.ID = userID
	c.SendCh = make(chan []byte, 256)
	c.Conn = conn

	defer close(c.SendCh)
	defer conn.Close()
	go c.writePump()
	c.readPump()
}
