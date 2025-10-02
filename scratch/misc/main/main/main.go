package main

import (
	"fmt"
	"math"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/life"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/when"
)

var signal = std.NewRevelation(Oscillator[float64](120, 0, 0.125))
var integral = std.NewRevelation[float64](Integrator[float64](signal))

func main() {
	cortex := std.NewCortex(std.RandomName())
	cortex.Frequency = 4 //hz

	cortex.Synapses() <- std.NewSynapse(life.Looping, "Record", func(imp *std.Impulse) {
		fmt.Println(signal.Reveal())
	}, nil)

	cortex.Synapses() <- std.NewSynapse(life.Looping, "Calculate", func(imp *std.Impulse) {
		integral.Reveal()
	}, when.Frequency(1.0))

	cortex.Spark()
	core.KeepAlive(time.Second * 3)
}

func Oscillator[TOut any](amplitude, phase, frequency any) func(time.Time) TOut {
	return OscillatorRef[TOut](&amplitude, &phase, &frequency)
}
func OscillatorRef[TOut any](amplitude, phase, frequency *any) func(time.Time) TOut {
	// TODO: Parse the input parameters to floats

	return func(time.Time) any {
		a := float64(*amplitude)
		p := float64(*phase)
		f := float64(*frequency)
		o := 2 * math.Pi * f

		moment := time.Now()
		t := moment.Sub(core.Inception).Seconds()
		value := a * math.Sin(o*t+p)

		// TODO: Parse the value to TOut
		return T(value)
	}
}

func Integrator[TOut any](signal any) func(time.Time) TOut {
	// TODO: Parse the revelation out of the signal

	return func(last time.Time) T {
		yield := signal.LatestSince(last)

		area := 0.0
		for i := len(yield) - 1; i > 0; i-- {
			dt := yield[i].Moment.Sub(yield[i-1].Moment).Seconds()
			if dt <= 0 {
				continue
			}
			area += dt * (float64(yield[i].Element) + float64(yield[i-1].Element)) * 0.5
		}

		last = time.Now()
		return T(area)
	}
}
