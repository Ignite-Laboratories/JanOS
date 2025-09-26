package main

import (
	"fmt"
	"sync"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/when"
)

func main() {
	c := std.NewCortex(std.RandomName())
	c.Frequency = 60 //hz

	freq := 2.0

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(&freq), Cleanup)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(&freq), Cleanup)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(&freq), Cleanup)

	c.Deferrals() <- func(wg *sync.WaitGroup) {
		time.Sleep(time.Second * 5)
		wg.Done()
	}

	c.Spark()
	core.KeepAlive()
}

func Cleanup(imp *std.Impulse, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	wg.Done()
}

var i = 0

func PrintCycle(imp *std.Impulse) {
	if i >= 3 {
		imp.Decay = true
	}
	i++

	fmt.Printf("\t%v [cycle] %v\n", imp.Timeline.CyclePeriod().String(), imp.Timeline.CyclePeriod().String())
}

func PrintRefractory(imp *std.Impulse) {
	fmt.Printf("\t%v [refraction] %v\n", imp.Timeline.CyclePeriod().String(), imp.Timeline.RefractoryPeriod().String())
}

func PrintResponse(imp *std.Impulse) {
	fmt.Printf("\t%v [response] %v\n", imp.Timeline.CyclePeriod().String(), imp.Timeline.ResponseTime().String())
}
