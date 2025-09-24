package main

import (
	"fmt"
	"time"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
)

func main() {
	// First graphical app: walk a lone dot across all pixels in a loop on every impulse of the engine
	// This should fire one frame per impulse, correlating to stepping one pixel

	c := std.NewCortex(std.RandomName())
	c.Frequency = 60
	//c.Mute()

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager A", Averager, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager B", Averager, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager C", Averager, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager D", Averager, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager E", Averager, nil)
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager F", Averager, nil)
	c.Spark()

	core.KeepAlive()
}

func Printer(imp *std.Impulse) bool {
	fmt.Println(imp.Timeline.CyclePeriod())
	return true
}

var averager time.Duration
var i int
var averageCap = 64

func Averager(imp *std.Impulse) bool {
	if imp.Timeline.Len() == 0 {
		return true
	}

	latest := imp.Timeline.LatestSince(time.Now().Add(-time.Second))
	fmt.Println(latest)

	averager += imp.Timeline.CyclePeriod()
	i++
	if i >= averageCap {
		fmt.Println(averager / time.Duration(i))
		averager = 0
		i = 0
	}
	time.Sleep(13 * time.Millisecond)
	return true
}
