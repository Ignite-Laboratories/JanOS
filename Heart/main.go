package main

import (
	"Common/RPC/PerceptionAPI"
	"github.com/google/uuid"
	"log"
)

var ThisComponent = PerceptionAPI.NewComponent(uuid.New().String(), "tcp", "localhost:421")

func main() {
	remote := ThisComponent.ConnectRemote("tcp", "localhost:420")

	// Look at the incoming packets
	for msg := range ThisComponent.Server.PacketChannel {
		log.Printf("[RPC] [%s] [Message] - %s", ThisComponent.Server.ID, msg)
		// Send them out the remote
		remote.ProcessPacket(msg)
	}
}
