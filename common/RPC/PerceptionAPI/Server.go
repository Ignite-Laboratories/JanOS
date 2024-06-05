package PerceptionAPI

import "Common/RPC"

type Server struct {
	Address       string
	PacketChannel chan string
}

func NewServer(address string) *Server {
	return &Server{
		Address:       address,
		PacketChannel: make(chan string),
	}
}

func (s *Server) Start() {
	h := RPC.NewHandler[API]("tcp", s.Address)
	h.API.Server = s
	h.StartServer()
}
