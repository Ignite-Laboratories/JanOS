package SampleAPI

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

func (a *Client) Echo(value string) string {
	var str string
	err := a.client.Call("API.Echo", value, &str)
	if err != nil {
		log.Fatal("[RPC] Error: ", err)
	}
	return str
}

func (a *Client) GetObject(title string) *TestObject {
	var obj *TestObject
	err := a.client.Call("API.GetObject", title, &obj)
	if err != nil {
		log.Fatal("[RPC] Error: ", err)
	}
	return obj
}

func (a *Client) ReceiveObject(obj *TestObject) {
	var str string
	err := a.client.Call("API.ReceiveObject", obj, &str)
	if err != nil {
		log.Fatal("[RPC] Error: ", err)
	}
}
