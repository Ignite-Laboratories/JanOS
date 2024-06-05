package TestAPI

import (
	"Common/RPC"
	"net/rpc"
)

type Client struct {
	client *rpc.Client
}

func NewClient(address string) *Client {
	h := new(RPC.StdRPCHandler[API])
	h.Network = "tcp"
	h.Address = address
	c := h.StartClient()

	return &Client{client: c}
}

func (a *Client) Echo(value string) string {
	var rstr string
	a.client.Call("API.Echo", "TEST", &rstr)
	return rstr
}

func (a *Client) GetObject(title string) *TestObject {
	var robj *TestObject
	a.client.Call("API.GetObject", "Test Object", &robj)
	return robj
}

func (a *Client) ReceiveObject(obj *TestObject) {
	a.client.Call("API.ReceiveObject", &obj, "")
}
