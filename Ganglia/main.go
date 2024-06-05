package main

import (
	"Common/RPC/PerceptionAPI"
	"log"
	"time"
)

func main() {
	//s := PerceptionAPI.NewServer("unix", "/tmp/perception.sock")
	s := PerceptionAPI.NewServer("tcp", "localhost:420")
	go s.Start()

	time.Sleep(time.Second)

	for msg := range s.PacketChannel {
		log.Println("[PACKET] " + msg)
	}
}
