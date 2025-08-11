package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/sys/atlas"
	"github.com/ignite-laboratories/core/sys/when"
	"github.com/ignite-laboratories/core/temporal"
)

func main() {
	var incrementer = temporal.Calculation(atlas.Impulse, when.Always, false, increment)
	temporal.Integration(atlas.Impulse, when.Frequency(std.Ref(1.0)), false, false, printTimeline, incrementer)
	atlas.Impulse.MaxFrequency = 4
	atlas.Impulse.Spark()
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
