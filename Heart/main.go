package main

import (
	"github.com/Ignite-Laboratories/JanOS/internal/state"
	"log"
)

func main() {
	s := state.Incept()
	remote := s.Nexus.ConnectRemote(s.Nexus.Server.PacketChannel, "Ganglia")

	// Look at the incoming packets
	for msg := range s.Nexus.Server.PacketChannel {
		log.Printf("[Backplane] [%s] [Message] - %s", s.Nexus.Server.ID, msg)
		// Send them out the remote
		remote.ProcessPacket(msg)
	}
}
