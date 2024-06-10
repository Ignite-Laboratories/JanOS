package nexus

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

func (h *Handler[Tapi]) StartClient() (*rpc.Client, error) {
	log.Printf("[nexus] Launching nexus client")

	client, err := rpc.DialHTTP(h.Network, h.Address)
	if err != nil {
		log.Println("[nexus] Connection error: ", err)
		return nil, err
	}

	return client, nil
}

func (h *Handler[Tapi]) StartServer() {
	for {
		err := h.startServer()
		if err != nil {
			log.Println("[nexus] Server failed erroneously", err)
		}
		log.Printf("[nexus] Restarting server [%s]\n", h.Address)
	}
}

func (h *Handler[Tapi]) startServer() error {
	log.Printf("[nexus] Launching nexus server")
	err := rpc.Register(h.API)
	if err != nil {
		log.Println("[nexus] Error registering API", err)
		return err
	}

	rpc.HandleHTTP()

	listener, err := net.Listen(h.Network, h.Address)
	if err != nil {
		log.Println("Error listening on "+h.Address, err)
		return err
	}
	log.Printf("[nexus] Listening on [%s]", h.Address)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Println("Error serving", err)
		return err
	}
	log.Printf("[nexus] Server Terminated")
	return nil
}
