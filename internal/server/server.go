package server

import (
	"SleepWalker/SystemMonitor/internal/client"
	"net/http"
	"sync"
)

type Server struct {
	MessageBuffer int
	mux           http.ServeMux
	mutex         sync.Mutex
	clients       map[*client.Client]struct{}
}

func NewServer() *Server {
	server := &Server{
		MessageBuffer: 10,
		clients:       make(map[*client.Client]struct{}),
	}
	server.mux.Handle("/", http.FileServer(http.Dir("./template")))
	server.mux.HandleFunc("/ws", server.handleWebSocket)
	return server
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {

}
