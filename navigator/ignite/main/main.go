package main

import (
	"os"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/std/neural"
	"git.ignitelabs.net/janos/core/sys/deploy"
	"git.ignitelabs.net/janos/navigator/ignite"
)

var port = "4242"
var cortex = std.NewCortex(std.RandomName())

func main() {
	if len(os.Args) > 1 && os.Args[1] == "deploy" {
		deploy.Fly.Spark("ignitelabs-net", "navigator", "ignite")
	} else {
		cortex.Frequency = 1 //hz
		cortex.Mute()

		cortex.Synapses() <- neural.Net.Server("ignitelabs.net", ":4242", ignite.Handler, func(imp *std.Impulse) {
			cortex.Impulse()
		})

		cortex.Spark()
		cortex.Impulse()
		core.KeepAlive()
	}
}
