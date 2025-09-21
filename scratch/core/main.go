package main

import (
	"fmt"
	"time"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
	"git.ignitelabs.net/navigator/netscape"
)

func main() {
	c := std.NewCortex(std.RandomName())
	c.Frequency = 1 //when.DurationToHertz[float64](3 * time.Second)

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Web Server", netscape.Neural.NavigateImpulse, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Key Watcher", KeyWatcher, nil)

	go func() {
		time.Sleep(5 * time.Second)
		c.Mute()
	}()

	c.Spark()
	core.KeepAlive()
}

func KeyWatcher(imp *std.Impulse) {
	fmt.Println(imp.Timeline.RefractoryPeriod())
}
