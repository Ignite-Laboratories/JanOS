package PerceptionAPI

import "Common/RPC"

type Server struct {
	Network       string
	Address       string
	PacketChannel chan string
}

func NewServer(network string, address string) *Server {
	return &Server{
		Network:       network,
		Address:       address,
		PacketChannel: make(chan string),
	}
}

func (s *Server) Start() {
	h := RPC.NewHandler[API](s.Network, s.Address)
	h.API.Server = s
	h.StartServer()
}
