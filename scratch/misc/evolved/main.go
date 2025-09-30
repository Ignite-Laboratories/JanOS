package main

import (
	"fmt"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/when"
)

/*
E0S0

This prints several statistics through phasing activation across time using two different techniques:

	- Cortical Beats, where the impulse.Beat value creates a manual round-robin effect
	- Clustering, where a neural cluster creates a round-robin effect

While cortical beats work, they come with several drawbacks.  Namely, their potentials must be manually created!
With a cortical cluster you can add and remove as many synapses as you like, and they will naturally be distributed
across a dynamically adjustable period of time.  This allows you to define a cluster of synapses which sequentially
activate at the same interval of time.  While you could add as many synapses as you'd like, their ability to completely
activate across a window of time is directly tied to the frequency of the driving cortex.  This can be seen by
setting the cortex frequency to a smaller value than the number of synapses in the cluster.

In addition, I've written two examples of performance considerations.  With chaotic synaptic activity you can clearly
see the processors pegged in a cyclic fashion.  With clustered activity, the amount of processor's load is almost
imperceptible - relatively speaking.
*/

var cortex = std.NewCortex(std.RandomName())

func main() {
	//UseCorticalBeat()
	//UseClustering()

	//ShowPerformantSynapticActivity()
	ShowImperformantSynapticActivity()

	cortex.Spark()
	core.KeepAlive()
}

func UseCorticalBeat() {
	cortex.Frequency = 2 //hz
	cortex.BeatPeriod = 2

	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, func(imp *std.Impulse) bool {
		return imp.Beat == 0
	})

	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, func(imp *std.Impulse) bool {
		return imp.Beat == 1
	})

	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, func(imp *std.Impulse) bool {
		return imp.Beat == 2
	})
}

func UseClustering() {
	cortex.Frequency = 120 //hz

	cluster := cortex.CreateCluster("cluster", core.Ref(time.Second), nil)

	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)

	cluster.Spark()
}

func ShowPerformantSynapticActivity() {
	cortex.Frequency = 11111 //hz

	cluster := cortex.CreateCluster("cluster", core.Ref(time.Second), nil)

	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Cycle Printer", PrintCycle, nil)
	cluster.Loop() <- std.NewNeuron("Refractory Printer", PrintRefractory, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)
	cluster.Loop() <- std.NewNeuron("Response Time Printer", PrintResponse, nil)

	cluster.Spark()
}

func ShowImperformantSynapticActivity() {
	cortex.Frequency = 11111 //hz

	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Cycle Printer", PrintCycle, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Refractory Printer", PrintRefractory, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
	cortex.Synapses() <- std.NewSynapse(lifecycle.Looping, "Response Time Printer", PrintResponse, when.Frequency(1))
}

func PrintCycle(imp *std.Impulse) {
	fmt.Printf("%v [cycle] %v\n", imp.Timeline.CyclePeriod().String(), imp.Timeline.CyclePeriod().String())
}

func PrintRefractory(imp *std.Impulse) {
	fmt.Printf("%v [refraction] %v\n", imp.Timeline.CyclePeriod().String(), imp.Timeline.RefractoryPeriod().String())
}

var decay = 0

func PrintResponse(imp *std.Impulse) {
	if decay == 2 {
		imp.Decay = true
	}
	//decay++
	fmt.Printf("%v [response] %v\n", imp.Timeline.CyclePeriod().String(), imp.Timeline.ResponseTime().String())
}
