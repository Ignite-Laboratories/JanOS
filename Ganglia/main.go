package main

import (
	"Common/RPC/TestAPI"
)

func main() {
	s := TestAPI.NewServer("localhost:420")
	s.Start()
}
