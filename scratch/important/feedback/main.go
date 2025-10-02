package main

import (
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/life"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/num"
	"git.ignitelabs.net/janos/core/sys/rec"
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
	cortex.Frequency = 11000 //hz

	cortex.Synapses() <- std.NewSynapse(life.Looping, "feedback loop", Print, Periodic(time.Second))

	cortex.Spark()
	core.KeepAlive()
}

var i = 0

func Print(imp *std.Impulse) {
	rec.Printf(imp.Bridge.String(), imp.Timeline.CyclePeriod().String()+"\n")
	if i > 5 {
		time.Sleep(time.Duration(num.RandomWithinRange[int](0, 2048)) * time.Millisecond)
	}
	if i > 10 {
		i = 0
	}
	i++
}

func Periodic(duration time.Duration) func(*std.Impulse) bool {
	last := time.Now()
	initial := true
	return func(imp *std.Impulse) bool {
		if initial {
			initial = false
			return true
		}
		now := time.Now()
		observed := now.Sub(last)
		adjustment := observed - duration

		if now.After(last.Add(duration).Add(-adjustment)) {
			last = now
			return true
		}
		return false
	}
}
