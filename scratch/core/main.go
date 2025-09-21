package main

import (
	"fmt"
	"time"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
	"git.ignitelabs.net/core/std/synapse"
	"git.ignitelabs.net/core/sys/given"
	"git.ignitelabs.net/core/sys/given/format"
	"git.ignitelabs.net/core/sys/when"
	"git.ignitelabs.net/navigator/netscape"
)

func main() {
	c := std.NewCortex(given.Random[format.Default]().Name)
	c.Frequency = when.DurationToHertz[float64](3 * time.Second)

	c.Synapses() <- synapse.NewLongRunning(lifecycle.Looping, "Web Server", netscape.Neural.NavigateImpulse, nil)
	c.Synapses() <- synapse.NewLongRunning(lifecycle.Looping, "Key Watcher", KeyWatcher, nil)

	c.Spark()
	core.KeepAlive()
}

func KeyWatcher(imp *std.Impulse) {
	fmt.Println("here")
}
