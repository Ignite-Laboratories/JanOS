package PerceptionAPI

import (
	"github.com/Ignite-Laboratories/JanOS/Internal/Backplane"
	"github.com/google/uuid"
	"log"
	"net/rpc"
)

type Client struct {
	ID      string
	handler *Backplane.Handler[API]
	client  *rpc.Client
}

func NewClient(network string, address string) *Client {
	return &Client{
		ID:      uuid.New().String(),
		handler: Backplane.NewHandler[API](network, address),
	}
}

func (a *Client) Start() {
	a.client = a.handler.StartClient()
}

func (a *Client) ProcessPacket(value string) {
	var str string
	err := a.client.Call("API.ProcessPacket", value, &str)
	if err != nil {
		log.Fatal("[Backplane] Error: ", err)
	}
}
