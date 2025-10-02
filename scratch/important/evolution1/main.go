package main

import (
	"important/evolution1/std"

	"git.ignitelabs.net/janos/core/enum/life"
	"git.ignitelabs.net/janos/core/sys/rec"
	"git.ignitelabs.net/janos/core/sys/when"
)

func main() {
	cortex := std.NewCortex("bob", when.HertzToDuration(1), 1)

	cortex <- std.Signal.Synapse(life.Looping, "Print", Printer, nil)

	cortex <- std.Signal.Spark()
	cortex.KeepAlive()

	t := Test{}
	t.Control <- nil
}

func Printer(imp *std.Impulse) {
	rec.Printf(imp.String(), imp.CyclePeriod().String()+"\n")
}

type Test struct {
	Control chan<- any

	Hello string
}
