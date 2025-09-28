package main

import (
	"os"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/std/neural"
	"git.ignitelabs.net/janos/core/sys/deploy"
)

var port = "4242"
var cortex = std.NewCortex(std.RandomName())

func main() {
	if len(os.Args) > 1 && os.Args[1] == "deploy" {
		deploy.Fly.Spark("git-ignitelabs-net", "navigator", "git")
	} else {
		cortex.Frequency = 1 //hz
		cortex.Mute()

		cortex.Synapses() <- neural.Net.Vanity("git.ignitelabs.net", "https://github.com/ignite-laboratories", 4242, func(imp *std.Impulse) {
			cortex.Impulse()
		})

		cortex.Spark()
		cortex.Impulse()
		core.KeepAlive()
	}
}
