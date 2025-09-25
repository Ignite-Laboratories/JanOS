package main

import (
	"fmt"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
)

func main() {
	c := std.NewCortex(std.RandomName())
	c.Frequency = 60 //hz

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, nil)

	c.Spark()
	core.KeepAlive()
}

func PrintCycle(imp *std.Impulse) bool {
	//fmt.Println(imp.Timeline.CyclePeriod())
	return true
}

func PrintRefractory(imp *std.Impulse) bool {
	fmt.Println(imp.Timeline.RefractoryPeriod())
	return true
}

func PrintResponse(imp *std.Impulse) bool {
	//fmt.Println(imp.Timeline.ResponseTime())
	return true
}
