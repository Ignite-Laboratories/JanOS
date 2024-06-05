package main

import (
	"Common/RPC"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	h := new(RPC.StdRPCHandler[RPC.TestAPI])
	h.Network = "tcp"
	h.Address = "localhost:420"
	c := h.StartClient()

	time.Sleep(time.Millisecond * 250)

	var rstr string
	c.Call("TestAPI.Echo", "TEST", &rstr)
	fmt.Println(rstr)

	var robj *RPC.TestObject
	c.Call("TestAPI.GetObject", "Test Object", &robj)
	js1, _ := json.Marshal(robj)
	fmt.Println(string(js1))

	c.Call("TestAPI.ReceiveObject", &robj, "")
	js2, _ := json.Marshal(robj)
	fmt.Println(string(js2))
}
