package main

import (
	"os"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/std/neural"
	"git.ignitelabs.net/janos/core/sys/rec"
)

func init() {
	rec.Verbose = false
}

func main() {
	var cortex *std.Cortex

	if len(os.Args) > 1 && os.Args[1] == "server" {
		core.Describe("Sub-Process")
		cortex = std.NewCortex(std.RandomName())

		cortex.Synapses() <- neural.Net.HelloWorld("Hello, World!", os.Args[2])
	} else {
		core.Describe("Multiplexer")
		cortex = std.NewCortex(std.RandomName())

		cortex.Synapses() <- neural.Shell.SubProcess("sub process a", []string{"go", "run", "./main", "server", ":4242"}, func(imp *std.Impulse) {
			cortex.Impulse()
		})
		cortex.Synapses() <- neural.Shell.SubProcess("sub process b", []string{"go", "run", "./main", "server", ":4243"}, func(imp *std.Impulse) {
			cortex.Impulse()
		})
		cortex.Synapses() <- neural.Shell.SubProcess("sub process c", []string{"go", "run", "./main", "server", ":4244"}, func(imp *std.Impulse) {
			cortex.Impulse()
		})
	}

	cortex.Frequency = 1 //hz
	cortex.Mute()
	cortex.Spark()
	cortex.Impulse()
	core.KeepAlive()
}
