package main

import (
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/enum/lifecycle"
	"git.ignitelabs.net/janos/core/std"
	"git.ignitelabs.net/janos/core/sys/rec"
)

func main() {
	c := std.NewCortex(std.RandomName())
	c.Frequency = 1 //hz

	c.Synapses() <- std.NewSynapse(lifecycle.Stimulative, "Print", Printer, nil)

	c.Spark()
	core.KeepAlive(time.Second * 5)
}

func Printer(imp *std.Impulse) {
	rec.Printf(imp.Bridge, "START: %v\n", time.Now())
	time.Sleep(time.Second * 3)
	rec.Printf(imp.Bridge, " STOP: %v\n", time.Now())
}
