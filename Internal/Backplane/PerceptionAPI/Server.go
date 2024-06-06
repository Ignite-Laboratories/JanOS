package PerceptionAPI

import (
	"github.com/Ignite-Laboratories/JanOS/Internal/Backplane"
	"github.com/Ignite-Laboratories/JanOS/Internal/Config"
)

type Server struct {
	ID            string
	Network       string
	Address       string
	PacketChannel chan string
}

func NewServer(network string, address string) *Server {
	return &Server{
		ID:            Config.Current.ID,
		Network:       network,
		Address:       address,
		PacketChannel: make(chan string),
	}
}

func (s *Server) Start() {
	h := Backplane.NewHandler[API](s.Network, s.Address)
	h.API.Server = s
	go h.StartServer()
}
