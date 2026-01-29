package main

import (
	"fmt"; "net/http"; "log"; "github.com/gorilla/websocket"; "html/template"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

func writeMessage(conn *websocket.Conn, outChannel chan []byte, msgType int) {
	for msg := range outChannel {
		if (msg[0] == 'q') {
			log.Printf("[WEBSOCKET] QUIT SIGNAL RECEIVED for writing channel")
			return
		}
		log.Printf("[WEBSOCKET] returning message: %s\n", msg)
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
		}
	}
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("[WEBSOCKET] connection received")
	outCh := make(chan []byte, 256)
	defer close(outCh)
	defer conn.Close()
	go writeMessage(conn, outCh, 1)
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

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ws", websocketHandler)
	log.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

