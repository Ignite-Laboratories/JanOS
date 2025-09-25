package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/beat"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"time"
)

var printPeriod = time.Second
var modulo = 1024
var observer = temporal.Observer(core.Impulse, beat.Modulo(&modulo), false, &core.Impulse.Last.Duration)

func main() {
	//observer.Resistance = 100000
	core.Impulse.Loop(KeepAlive, when.Always, false)
	core.Impulse.Loop(PrintAnalysis, when.Duration(&printPeriod), false)
	core.Impulse.Loop(Toggle, when.Always, false)
	core.Impulse.Spark()
}

func Toggle(ctx core.Context) {
	if observer.Stimulator.Muted {
		fmt.Println("Unmuting")
		observer.Unmute()
	} else {
		fmt.Println("Muting")
		observer.Mute()
	}
	time.Sleep(time.Second * 7)
}

func KeepAlive(ctx core.Context) {
	for core.Alive {
		// This just lets the beat number count up infinitely for modulo use
	}
}

func PrintAnalysis(ctx core.Context) {
	observer.Mutex.Lock()
	data := make([]std.Data[time.Duration], len(observer.Timeline))
	copy(data, observer.Timeline)
	observer.Mutex.Unlock()
	out := make([]time.Duration, len(data))
	for i, d := range data {
		out[i] = d.Point
	}
	fmt.Println(out)
}
