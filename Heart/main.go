package main

import (
	"github.com/Ignite-Laboratories/JanOS/Internal/Component"
	"log"
)

func main() {
	Component.Setup()
	remote := Component.This.ConnectRemote(Component.This.Server.PacketChannel, "tcp", "127.0.0.1:421")

	// Look at the incoming packets
	for msg := range Component.This.Server.PacketChannel {
		log.Printf("[Backplane] [%s] [Message] - %s", Component.This.Server.ID, msg)
		// Send them out the remote
		remote.ProcessPacket(msg)
	}
}
