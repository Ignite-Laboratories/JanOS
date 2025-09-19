package main

import (
	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
	"git.ignitelabs.net/core/sys/given"
	"git.ignitelabs.net/core/sys/given/format"
	"git.ignitelabs.net/core/sys/log"
)

func main() {
	defer func() {
		core.ShutdownNow()
	}()

	c := std.NewCortex(given.Random[format.Default]().Name)

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, neuron{
		Entity: std.NewEntity[format.Default](),
	})

	c.Spark()
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
