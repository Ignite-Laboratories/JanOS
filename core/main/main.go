package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/direction/ordinal"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/pad/pad4d"
	"github.com/ignite-laboratories/core/sys/pad/scheme"
)

func main() {
	seen := make(map[int16]struct{})
	limit := 0

	fmt.Println(pad4d.Align(scheme.Reverse, ordinal.Positive, 10, [][][][]int16{}, func() int16 {
		keepGoing := true
		var rnd int16
		for keepGoing {
			if limit > 44 {
				limit = 0
				seen = make(map[int16]struct{})
			}
			rnd = num.RandomWithinRange[int16](0, 44)
			if _, ok := seen[rnd]; !ok {
				keepGoing = false
				seen[rnd] = struct{}{}
				limit++
			}
		}
		return rnd
	}))
}
