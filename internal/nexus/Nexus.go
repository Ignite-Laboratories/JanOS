package nexus

import (
	"github.com/Ignite-Laboratories/JanOS/internal/common"
	"log"
)

type Nexus struct {
	ID          string
	Name        string
	Title       string
	Description string
	Clients     map[string]*Client
	Server      *Server
}

func NewNexus() *Nexus {
	log.Printf("[nexus] Initializing NexusAPI [%s]\n", common.GetComponentName())
	c := Nexus{
		ID:          common.Current.ID,
		Name:        common.GetComponentName(),
		Title:       common.Current.Title,
		Description: common.Current.Description,
		Clients:     make(map[string]*Client),
		Server:      NewServer("unix", common.GetSocketAddress()),
	}

	c.Server.Start()

	return &c
}

func (c *Nexus) ConnectRemote(channel chan string, componentName string) *Client {
	return connectRemote(c, channel, "unix", common.GetRemoteSocketAddress(componentName))
}

func (c *Nexus) ConnectRemoteInstance(channel chan string, componentName string, instanceName string) *Client {
	return connectRemote(c, channel, "unix", common.GetRemoteSocketAddress(componentName, instanceName))
}

func (c *Nexus) ConnectTCPRemote(channel chan string, address string) *Client {
	return connectRemote(c, channel, "tcp", address)
}

func connectRemote(c *Nexus, channel chan string, network string, address string) *Client {
	log.Printf("[nexus] Connecting to remote [%s]\n", address)
	client := NewClient(network, address)
	client.Start(channel)

	c.Clients[client.ID] = client
	log.Printf("[nexus] Remote Connected [%s]\n", client.ID)
	return client
}
