package SampleAPI

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
	RPC.NewHandler[API]("tcp", s.Address).StartServer()
}
