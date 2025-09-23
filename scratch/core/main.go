package main

import (
	"fmt"
	"time"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
)

func main() {
	c := std.NewCortex(std.RandomName())
	c.Frequency = 60
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager", Averager, nil)
	//c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Printer", Printer, nil)
	c.Spark()

	core.KeepAlive()
}

func Printer(imp *std.Impulse) {
	if imp.Timeline.Last == nil {
		return
	}
	fmt.Println(imp.Timeline.CyclePeriod())
	time.Sleep(7 * time.Millisecond)
}

var averager time.Duration
var i int
var averageCap = 64

func Averager(imp *std.Impulse) {
	if imp.Timeline.Last == nil {
		return
	}
	averager += imp.Timeline.CyclePeriod()
	i++
	if i >= averageCap {
		fmt.Println(averager / time.Duration(i))
		averager = 0
		i = 0
	}
	time.Sleep(13 * time.Millisecond)
}
