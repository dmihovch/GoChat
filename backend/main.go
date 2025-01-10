package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.webSocketHandler))
	http.ListenAndServe(":3000", nil)
}
