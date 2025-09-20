package main

import (
	"fmt"
	"time"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
	"git.ignitelabs.net/core/sys/log"
)

func main2() {
	printer := make(chan time.Duration, 1<<16)

	go func() {
		for p := range printer {
			fmt.Println(p)
		}
	}()

	last := time.Now()
	for core.Alive() {
		now := time.Now()
		printer <- now.Sub(last)
		last = now
	}
}

func main() {
	n := neuron{
		Entity: std.NewEntityNamed("Bartholomew"),
	}

	createCortex("A", n)
	createCortex("B", n)
	createCortex("C", n)
	createCortex("D", n)
	createCortex("E", n)

	core.KeepAlive()
}

func createCortex(cortexName string, n neuron) {
	c := std.NewCortex(cortexName)

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
