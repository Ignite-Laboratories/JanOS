package main

import (
	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
	"git.ignitelabs.net/core/sys/log"
)

func main() {
	n := neuron{
		Entity: std.NewEntityNamed("Bartholomew"),
	}

	createCortex("A", n)
	createCortex("B", n)
	createCortex("C", n)
	createCortex("D", n)
	createCortex("E", n)
	createCortex("F", n)
	createCortex("G", n)
	createCortex("H", n)
	createCortex("I", n)
	createCortex("J", n)
	createCortex("K", n)
	createCortex("L", n)
	createCortex("M", n)
	createCortex("N", n)
	createCortex("O", n)
	createCortex("P", n)
	createCortex("Q", n)
	createCortex("R", n)
	createCortex("S", n)
	createCortex("T", n)
	createCortex("U", n)
	createCortex("V", n)
	createCortex("W", n)
	createCortex("X", n)
	createCortex("Y", n)
	createCortex("Z", n)

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

func (neuron) Action(imp std.Impulse) {
	log.Printf(imp.Bridge, "%v\n", imp.Timeline.ResponseTime())
}

func (neuron) Potential(ctx std.Impulse) bool {
	return true
}

func (neuron) Cleanup() {

}
