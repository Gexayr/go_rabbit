package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type Websocket struct {
	clients    map[net.Conn]bool
	register   chan net.Conn
	unregister chan net.Conn
}

func NewWebsocket() *Websocket {
	return &Websocket{
		clients:    make(map[net.Conn]bool),
		register:   make(chan net.Conn),
		unregister: make(chan net.Conn),
	}
}

func (ws *Websocket) handleConnections() {
	for {
		select {
		case conn := <-ws.register:
			ws.clients[conn] = true
		case conn := <-ws.unregister:
			if _, ok := ws.clients[conn]; ok {
				conn.Close()
				delete(ws.clients, conn)
			}
		}
	}
}

func (ws *Websocket) serveWs(w http.ResponseWriter, r *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	conn, _, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	ws.register <- conn
	defer func() { ws.unregister <- conn }()

	for {
		message := make([]byte, 1024)
		_, err := conn.Read(message)
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				return
			}
			log.Println(err)
			return
		}
		// Обработка полученного сообщения
	}
}

func main() {
	websocket := NewWebsocket()
	go websocket.handleConnections()

	http.HandleFunc("/ws", websocket.serveWs)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
