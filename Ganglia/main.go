package main

import (
	"Common/RPC/PerceptionAPI"
	"github.com/google/uuid"
	"log"
	"time"
)

var ThisComponent = PerceptionAPI.NewComponent(uuid.New().String(), "tcp", "localhost:420")

func main() {
	Kickstart()
}

func Kickstart() {
	time.Sleep(time.Second * 5)

	remote := ThisComponent.ConnectRemote("tcp", "localhost:421")
	remote.ProcessPacket("HELLO!")

	// Look at the incoming packets
	for msg := range ThisComponent.Server.PacketChannel {
		log.Printf("[RPC] [%s] [Message] - %s", ThisComponent.Server.ID, msg)
		// Send them out the remote
		remote.ProcessPacket(msg)
	}
}
