package PerceptionAPI

import (
	"log"
)

type Component struct {
	ID      string
	Clients map[string]*Client
	Server  *Server
}

func NewComponent(id string, network string, address string) *Component {
	log.Printf("[RPC] Initializing Component [%s]\n", address)
	c := Component{
		ID:      id,
		Clients: make(map[string]*Client),
		Server:  NewServer(network, address),
	}

	c.Server.Start()

	return &c
}

func (c *Component) ConnectRemote(network string, address string) *Client {
	log.Printf("[RPC] Connecting to remote [%s]\n", address)
	client := NewClient(network, address)
	client.Start()

	c.Clients[client.ID] = client
	log.Printf("[RPC] Remote Connected [%s]\n", client.ID)
	return client
}
