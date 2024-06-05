package TestAPI

import (
	"Common/RPC"
	"log"
	"net/rpc"
)

type Client struct {
	client *rpc.Client
}

func NewClient(address string) *Client {
	h := new(RPC.Handler[API])
	h.Network = "tcp"
	h.Address = address
	c := h.StartClient()

	return &Client{client: c}
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
