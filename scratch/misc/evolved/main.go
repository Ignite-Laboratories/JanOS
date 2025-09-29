package main

import (
	"fmt"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/when"
)

/*
E0S0

This prints several statistics through phasing activation across time using two different techniques:

	- Cortical Phasing, where the impulse.Beat value creates a manual round-robin effect
	- Clustering, where a neural cluster creates a stable and safe round-robin effect

The reason cortical phasing is unstable is that neurons can decay, causing delayed gaps in the impulse cycle
which don't get "refilled."  When using a cluster this issue is resolved and decayed neurons naturally remove
from the cycle.
*/

var cortex = std.NewCortex(std.RandomName())

func main() {
	cortex.Frequency = 2 //hz

	//UseCorticalPhasing()
	UseClustering()

	cortex.Spark()
	core.KeepAlive()
}

func UseCorticalPhasing() {
	cortex.Phase = 2

	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, func(imp *std.Impulse) bool {
		return imp.Beat == 0
	})

	decay := 0
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, func(imp *std.Impulse) bool {
		if imp.Beat == 1 {
			decay++
			if decay == 2 {
				imp.Decay = true
			}
			return !imp.Decay
		}
		return false
	})

	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, func(imp *std.Impulse) bool {
		return imp.Beat == 2
	})
}

func UseClustering() {
	// 0 - Create the synaptic cluster
	syn, cluster := std.NewSynapticCluster("cluster", cortex, when.Frequency(8))

	// 1 - Add it to the cortex
	cortex.Synapses() <- syn

	// 2 - Add neurons to the cluster
	cluster <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)

	// 3 - Add a faux delayed decay to the final neuron for demonstration purposes
	decay := 0
	cluster <- std.NewNeuron("Response Time Printer", PrintResponse, func(imp *std.Impulse) bool {
		if decay == 2 {
			imp.Decay = true
		}
		decay++
		return !imp.Decay
	})
}

func PrintCycle(imp *std.Impulse) {
	fmt.Printf("%v [cycle] %v\n", imp.Timeline.CyclePeriod().String(), imp.Timeline.CyclePeriod().String())
}

func PrintRefractory(imp *std.Impulse) {
	fmt.Printf("%v [refraction] %v\n", imp.Timeline.CyclePeriod().String(), imp.Timeline.RefractoryPeriod().String())
}

func PrintResponse(imp *std.Impulse) {
	fmt.Printf("%v [response] %v\n", imp.Timeline.CyclePeriod().String(), imp.Timeline.ResponseTime().String())
}
