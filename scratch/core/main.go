package main

import (
	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
	"git.ignitelabs.net/core/sys/log"
)

func main() {
	n := neuron{
		Entity: std.NewEntityNamed("Bartholemew"),
	}

	createCortex("A", n)
	createCortex("B", n)
	createCortex("C", n)

	core.KeepAlive()
}

func createCortex(cortexName string, n neuron) {
	c := std.NewCortex(cortexName)
	c.Frequency = 1

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, n)

	go c.Spark()
}

type neuron struct {
	std.Entity
}

func (neuron) Mute() {

}

func (neuron) Unmute() {

}

func (neuron) Action(ctx std.Context) {
	log.Printf(ctx.ModuleName, "Action\n")
}

func (neuron) Potential(ctx std.Context) bool {
	return true
}

func (neuron) Cleanup() {

}
