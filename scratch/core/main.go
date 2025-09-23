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
	c.Frequency = 240 //when.DurationToHertz[float64](3 * time.Minute)

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Web Server", netscape.Neural.NavigateImpulse, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Key Watcher", KeyWatcher, nil)

	c.Spark()
	core.KeepAlive()
}

var averager time.Duration
var i int
var averageCap = 64

func KeyWatcher(imp *std.Impulse) {
	if imp.Timeline.Last == nil {
		return
	}
	averager += imp.Timeline.Inception.Sub(imp.Timeline.Last.Inception)
	i++
	if i >= averageCap {
		fmt.Println(averager / time.Duration(i))
		averager = 0
		i = 0
	}
}
