package main

import (
	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/std/neural"
)

func main() {
	core.Describe("A 'Hello, World!' Server")

	cortex := std.NewCortex(std.RandomName())
	cortex.Frequency = 1 //hz
	cortex.Mute()

	cortex.Synapses() <- neural.Shell.SubProcess(lifecycle.Triggered, "Echo", []string{"echo", "Hello, World!"})

	cortex.Spark()
	cortex.Impulse()
	core.KeepAlive()
}
