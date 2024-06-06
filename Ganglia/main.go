package main

import (
	"github.com/Ignite-Laboratories/JanOS/Internal/Component"
	"log"
	"time"
)

func main() {
	Component.Setup()
	Kickstart()
}

func Kickstart() {
	time.Sleep(time.Second * 5)

	remote := Component.This.ConnectRemote("tcp", "127.0.0.1:422")
	remote.ProcessPacket("HELLO!")

	// Look at the incoming packets
	for msg := range Component.This.Server.PacketChannel {
		log.Printf("[Backplane] [%s] [Message] - %s", Component.This.Server.ID, msg)
		// Send them out the remote
		remote.ProcessPacket(msg)
	}
}
