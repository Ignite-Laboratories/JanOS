package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
	"math"
)

func main() {
	a := num.Compare(math.Inf(1), math.Inf(-1))
	fmt.Println(a)
}
