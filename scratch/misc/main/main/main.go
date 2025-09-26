package main

import (
	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/rec"
	"git.ignitelabs.net/janos/core/sys/when"
)

func main() {
	c := std.NewCortex(std.RandomName())
	c.Frequency = 1 //hz

	rec.Printf(c.Named(), "Hello, World!\n")

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Print", Printer, when.FrequencyRef(&frequency))

	c.Spark()
	core.KeepAlive()
}

var frequency = 1.0
var toggle = true

func Printer(imp *std.Impulse) {
	if toggle {
		frequency = 0.5
	} else {
		frequency = 1.0
	}
	toggle = !toggle
	rec.Printf(imp.Bridge, "%v\n", imp.Timeline.CyclePeriod())
}
