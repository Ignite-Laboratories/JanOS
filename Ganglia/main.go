package main

import (
	"Common/RPC/PerceptionAPI"
	"log"
	"time"
)

func main() {
	s := PerceptionAPI.NewServer("localhost:420")
	go s.Start()

	time.Sleep(time.Second)

	for msg := range s.PacketChannel {
		log.Println("[PACKET] " + msg)
	}
}

func HandlePackets(s *PerceptionAPI.Server) {
	for msg := range s.PacketChannel {
		log.Println("[PACKET] " + msg)
	}
}
