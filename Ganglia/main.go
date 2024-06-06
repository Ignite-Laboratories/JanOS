package main

import (
	"github.com/Ignite-Laboratories/JanOS/Internal/Component"
	"time"
)

func main() {
	Component.Setup()
	Kickstart()
	for {

	}
}

func Kickstart() {
	remote := Component.This.ConnectRemote(Component.This.Server.PacketChannel, "tcp", "127.0.0.1:422")
	for {
		remote.ProcessPacket("HELLO!")
		time.Sleep(time.Second)
	}
}
