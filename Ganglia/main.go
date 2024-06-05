package main

import (
	"Common/RPC/PerceptionAPI"
	"github.com/google/uuid"
	"log"
)

var ThisComponent = PerceptionAPI.NewComponent(uuid.New().String(), "tcp", "localhost:420")

func main() {
	for msg := range ThisComponent.Server.PacketChannel {
		log.Println("[PACKET] " + msg)
	}
}
