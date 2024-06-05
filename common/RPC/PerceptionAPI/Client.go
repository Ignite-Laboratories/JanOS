package PerceptionAPI

import (
	"Common/RPC"
	"log"
	"net/rpc"
)

type Client struct {
	handler *RPC.Handler[API]
	client  *rpc.Client
}

func NewClient(address string) *Client {
	return &Client{
		handler: RPC.NewHandler[API]("tcp", address),
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
