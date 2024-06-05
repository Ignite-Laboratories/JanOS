package RPC

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
	log.Printf("[RPC] Launching RPC client")

	client, err := rpc.DialHTTP(h.Network, h.Address)
	if err != nil {
		log.Fatal("[RPC] Connection error: ", err)
	}

	return client
}

func (h *Handler[Tapi]) StartServer() {
	log.Printf("[RPC] Launching RPC server")
	err := rpc.Register(h.API)
	if err != nil {
		log.Fatal("[RPC] Error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen(h.Network, h.Address)
	if err != nil {
		log.Fatal("Error listening on "+h.Address, err)
	}
	log.Printf("[RPC] Listening on [%s]", h.Address)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving", err)
	}
	log.Printf("[RPC] Server Terminated")
}
