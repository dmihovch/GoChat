package main

import (
	"fmt"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
	mutex sync.Mutex
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) addConnection(ws *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.conns[ws] = true
	fmt.Println("New Connection:", ws.RemoteAddr())
}
func (s *Server) removeConnection(ws *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.conns, ws)
	ws.Close()
	fmt.Println("Connection Removed:", ws.RemoteAddr())
}

func (s *Server) readMessages(ws *websocket.Conn) {
	buffer := make([]byte, 1024)
	for {
		numBytes, err := ws.Read(buffer)
		if err != nil {
			fmt.Println("Read Error:", err)
			s.removeConnection(ws)
			break
		}
		message := buffer[:numBytes]
		fmt.Println("Message from Client:", string(message))
		s.broadcastMessage(message)
	}
}

func (s *Server) broadcastMessage(msg []byte) {
	for socket := range s.conns {
		go func(socket *websocket.Conn) {
			if _, err := socket.Write(msg); err != nil {
				fmt.Println("Write error:", err)
				s.removeConnection(socket)
			}
		}(socket)
	}
}

func (s *Server) webSocketHandler(ws *websocket.Conn) {
	s.addConnection(ws)
	s.readMessages(ws)
}
