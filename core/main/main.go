package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
)

func main() {
	//fmt.Println(num.Smallest[int](-123.5, int8(-5)))
	//fmt.Println(num.Smallest[float64](-123.5, int8(5)))
	fmt.Println(num.Smallest[num.Morsel](123.5, -5))
	//fmt.Println(num.Smallest[int](123.5, int8(5)))
	//fmt.Println(num.Smallest[int](-0.0, 0.0))
}
