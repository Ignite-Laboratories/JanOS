package nexus

import (
	"github.com/Ignite-Laboratories/JanOS/internal/common"
)

type Server struct {
	ID            string
	Network       string
	Address       string
	PacketChannel chan string
}

func NewServer(network string, address string) *Server {
	return &Server{
		ID:            common.Current.ID,
		Network:       network,
		Address:       address,
		PacketChannel: make(chan string),
	}
}

func (s *Server) Start() {
	h := NewHandler[API](s.Network, s.Address)
	h.API.Server = s
	go h.StartServer()
}
