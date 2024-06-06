package SampleAPI

import "github.com/Ignite-Laboratories/JanOS/Internal/Backplane"

type Server struct {
	Address string
}

func NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

func (s *Server) Start() {
	Backplane.NewHandler[API]("tcp", s.Address).StartServer()
}
