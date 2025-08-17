package main

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
)

func main() {
	for {
		b, _ := bounded.By[float64](0, -2, 4)
		a := std.UniqueSet[float64]{
			Bounded: b,
		}

		for {
			a.Random(50)
		}

	}
}
