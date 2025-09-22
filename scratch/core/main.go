package main

import (
	"fmt"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
	"git.ignitelabs.net/navigator/netscape"
)

func main() {
	c := std.NewCortex(std.RandomName())
	c.Frequency = 60 //when.DurationToHertz[float64](3 * time.Second)

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Web Server", netscape.Neural.NavigateImpulse, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Key Watcher", KeyWatcher, nil)

	c.Spark()
	core.KeepAlive()
}

func KeyWatcher(imp *std.Impulse) {
	if imp.Timeline.Last == nil {
		return
	}

	fmt.Println(imp.Timeline.Inception.Sub(imp.Timeline.Last.Inception))
}
