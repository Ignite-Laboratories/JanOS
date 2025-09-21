package synapse

import (
	"sync"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
	"git.ignitelabs.net/core/std/neuron"
	"git.ignitelabs.net/core/sys/log"
)

// NewLongRunning creates a synaptic bridge to a long-running neuron, such as a blocking process (like a web server).  You may optionally provide
// 'nil' to the potential if you'd like to imply 'always fire'.
func NewLongRunning(lifecycle lifecycle.Lifecycle, neuronName string, action func(*std.Impulse), potential func(*std.Impulse) bool, cleanup ...func(*std.Impulse, *sync.WaitGroup)) std.Synapse {
	n := neuron.NewLongRunning(neuronName, action, potential, cleanup...)
	log.Printf(core.ModuleName, "created long-running neuron '%s'\n", n.Named())
	return std.NewSynapse(lifecycle, n)
}
