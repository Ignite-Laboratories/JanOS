package state

import (
	"github.com/Ignite-Laboratories/JanOS/internal/common"
	"github.com/Ignite-Laboratories/JanOS/internal/nexus"
)

type State struct {
	Nexus *nexus.Nexus
}

func Incept() *State {
	// Setup config.Current from the provided information
	common.Initialize()

	// Setup the message nexus
	n := nexus.NewNexus()

	// Return the state machine
	return &State{
		Nexus: n,
	}
}
