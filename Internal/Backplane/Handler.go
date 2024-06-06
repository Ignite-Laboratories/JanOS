package Backplane

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Handler[Tapi any] struct {
	API     *Tapi
	Network string
	Address string
}

func NewHandler[Tapi any](network, address string) *Handler[Tapi] {
	return &Handler[Tapi]{
		API:     new(Tapi),
		Network: network,
		Address: address,
	}
}

func (h *Handler[Tapi]) StartClient() *rpc.Client {
	log.Printf("[Backplane] Launching Backplane client")

	client, err := rpc.DialHTTP(h.Network, h.Address)
	if err != nil {
		log.Fatal("[Backplane] Connection error: ", err)
	}

	return client
}

func (h *Handler[Tapi]) StartServer() {
	log.Printf("[Backplane] Launching Backplane server")
	err := rpc.Register(h.API)
	if err != nil {
		log.Fatal("[Backplane] Error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen(h.Network, h.Address)
	if err != nil {
		log.Fatal("Error listening on "+h.Address, err)
	}
	log.Printf("[Backplane] Listening on [%s]", h.Address)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving", err)
	}
	log.Printf("[Backplane] Server Terminated")
}
