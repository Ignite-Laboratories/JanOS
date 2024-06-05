package PerceptionAPI

import (
	"Common/RPC"
	"github.com/google/uuid"
)

type Server struct {
	ID            string
	Network       string
	Address       string
	PacketChannel chan string
}

func NewServer(network string, address string) *Server {
	return &Server{
		ID:            uuid.New().String(),
		Network:       network,
		Address:       address,
		PacketChannel: make(chan string),
	}
}

func (s *Server) Start() {
	h := RPC.NewHandler[API](s.Network, s.Address)
	h.API.Server = s
	go h.StartServer()
}
