package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
)

func main() {
	var incrementer = temporal.Calculation(core.Impulse, when.Always, false, increment)
	temporal.Integration(core.Impulse, when.Frequency(std.Ref(1.0)), false, false, printTimeline, incrementer)
	core.Impulse.MaxFrequency = 4
	core.Impulse.Spark()
}

var value = 0

func increment(ctx core.Context) int {
	value++
	return value
}

func printTimeline(ctx core.Context, cache *int, data []std.Data[int]) int {
	total := 0
	values := make([]int, len(data))
	for i, v := range data {
		values[i] = v.Point
		total += v.Point
	}
	*cache += total

	// Print the stats
	fmt.Printf("%v - %d - %d\n", values, total, *cache)
	return total
}
