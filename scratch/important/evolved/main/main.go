package main

import (
	"fmt"
	"math"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/life"
	"git.ignitelabs.net/janos/core/std"
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
	cortex.Frequency = 60 //hz

	cortex.Synapses() <- std.NewSynapse(life.Looping, "printer", func(imp *std.Impulse) {
		fmt.Println(cortex.PhaseAngle(0.25))
	}, nil)

	cortex.Spark()
	core.KeepAlive()
}

type PhaseAcc struct {
	Theta float64   // accumulated angle in [0, 2π)
	Last  time.Time // last update time
}

func (p *PhaseAcc) Init(now time.Time, initialPhase float64) {
	p.Last = now
	p.Theta = wrap2Pi(initialPhase)
}

func (p *PhaseAcc) Update(now time.Time, freqHz float64) {
	dt := now.Sub(p.Last).Seconds()
	if dt <= 0 {
		return
	}
	p.Theta = wrap2Pi(p.Theta + 2*math.Pi*freqHz*dt)
	p.Last = now
}

func (p *PhaseAcc) Angle() float64    { return p.Theta }
func (p *PhaseAcc) Fraction() float64 { return p.Theta / (2 * math.Pi) }
func (p *PhaseAcc) Sine() float64     { return math.Sin(p.Theta) }

func wrap2Pi(x float64) float64 {
	x = math.Mod(x, 2*math.Pi)
	if x < 0 {
		x += 2 * math.Pi
	}
	return x
}
