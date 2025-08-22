package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/pad"
	"math"
	"strconv"
)

func main() {
	a := strconv.FormatFloat(math.Inf(1), 'f', -1, 64)
	b := strconv.FormatFloat(math.Inf(-1), 'f', -1, 64)
	fmt.Println(a, b)

	//fmt.Println(num.Smallest[int](-123.5, int8(-5)))
	//fmt.Println(num.Smallest[float64](-123.5, int8(5)))
	//fmt.Println(num.Smallest[num.Morsel](123.5, -5))
	//fmt.Println(num.Smallest[int](123.5, int8(5)))
	//fmt.Println(num.Smallest[int](-0.0, 0.0))
	for {
		fmt.Println(pad.Any1D[num.Morsel, orthogonal.Left](33, num.RandomSetWithinRange(15, num.Morsel(5), 10), []num.Morsel{2}))
	}
}
