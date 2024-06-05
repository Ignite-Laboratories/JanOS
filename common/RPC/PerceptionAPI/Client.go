package PerceptionAPI

import (
	"Common/RPC"
	"github.com/google/uuid"
	"log"
	"net/rpc"
)

type Client struct {
	ID      string
	handler *RPC.Handler[API]
	client  *rpc.Client
}

func NewClient(network string, address string) *Client {
	return &Client{
		ID:      uuid.New().String(),
		handler: RPC.NewHandler[API](network, address),
	}
}

func (a *Client) Start() {
	a.client = a.handler.StartClient()
}

func (a *Client) ProcessPacket(value string) {
	var str string
	err := a.client.Call("API.ProcessPacket", value, &str)
	if err != nil {
		log.Fatal("[RPC] Error: ", err)
	}
}
