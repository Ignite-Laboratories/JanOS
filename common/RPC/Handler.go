package RPC

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Handler[Tapi any] struct {
	Network string
	Address string
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
	api := new(Tapi)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("[RPC] Error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen(h.Network, h.Address)
	if err != nil {
		log.Fatal("Error listening on "+h.Address, err)
	}
	log.Printf("[RPC] Listening on " + h.Address)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving", err)
	}
	log.Printf("[RPC] Serving on " + h.Address)
}
