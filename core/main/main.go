package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/vector"
)

func main() {
	a := vector.From7D[float64](1, 2, 3, 4, 5, 6, 7)
	fmt.Println(a)
}
