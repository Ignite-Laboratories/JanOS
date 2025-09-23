package main

import (
	"fmt"
	"sync"
	"time"

	"git.ignitelabs.net/core"
	"git.ignitelabs.net/core/enum/lifecycle"
	"git.ignitelabs.net/core/std"
)

func main() {
	c := std.NewCortex(std.RandomName())
	c.Frequency = 60
	c.Mute()

	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager A", Averager, nil)
	c.Deferrals() <- func(wg *sync.WaitGroup) {
		time.Sleep(time.Second)
		fmt.Println("A")
		wg.Done()
	}
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager B", Averager, nil)
	c.Deferrals() <- func(wg *sync.WaitGroup) {
		time.Sleep(time.Second)
		fmt.Println("B")
		wg.Done()
	}
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager C", Averager, nil)
	c.Deferrals() <- func(wg *sync.WaitGroup) {
		time.Sleep(time.Second)
		fmt.Println("C")
		wg.Done()
	}
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager D", Averager, nil)
	c.Deferrals() <- func(wg *sync.WaitGroup) {
		time.Sleep(time.Second)
		fmt.Println("D")
		wg.Done()
	}
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager E", Averager, nil)
	c.Deferrals() <- func(wg *sync.WaitGroup) {
		time.Sleep(time.Second)
		fmt.Println("E")
		wg.Done()
	}
	c.Synapses() <- std.NewSynapse(lifecycle.Looping, "Impulse Averager F", Averager, nil)
	c.Deferrals() <- func(wg *sync.WaitGroup) {
		time.Sleep(time.Second)
		fmt.Println("F")
		wg.Done()
	}
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
var averageCap = 1

var x int

func Averager(imp *std.Impulse) bool {
	if imp.Timeline.Last == nil {
		return true
	}
	averager += imp.Timeline.CyclePeriod()
	i++
	if i >= averageCap {
		fmt.Println(averager / time.Duration(i))
		averager = 0
		i = 0
	}
	time.Sleep(13 * time.Millisecond)
	x++
	return true

	if x >= 11*averageCap {
		return false
	}
	return true
}
