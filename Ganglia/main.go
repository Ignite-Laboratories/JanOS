package main

import (
	"github.com/Ignite-Laboratories/JanOS/internal/state"
	"time"
)

func main() {
	s := state.Incept()
	Kickstart(s)
	select {}
}

func Kickstart(s *state.State) {
	remote := s.Nexus.ConnectRemote(s.Nexus.Server.PacketChannel, "heart")
	for {
		remote.ProcessPacket("HELLO!")
		time.Sleep(time.Second)
	}
}
