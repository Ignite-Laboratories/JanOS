package main

import "Common/RPC"

func main() {
	h := new(RPC.StdRPCHandler[RPC.TestAPI])
	h.Network = "tcp"
	h.Address = "localhost:420"
	h.StartServer()
}
