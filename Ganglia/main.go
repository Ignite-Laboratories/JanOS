package main

import (
	"Common/RPC/SampleAPI"
)

func main() {
	s := SampleAPI.NewServer("localhost:420")
	s.Start()
}
