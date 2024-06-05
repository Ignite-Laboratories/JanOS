package TestAPI

import "Common/RPC"

type Server struct {
	Address string
}

func NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

func (s *Server) Start() {
	h := new(RPC.Handler[API])
	h.Network = "tcp"
	h.Address = s.Address
	h.StartServer()
}
