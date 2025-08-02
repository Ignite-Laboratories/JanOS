package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
)

// This example demonstrates how to use a custom random number generation,
// as well as normalization and re-scaling the generic numbers.

func main() {
	core.SetRandomNumberGenerator[byte](BadGenerator)

	for i := 0; i < 3; i++ {
		c := std.RandomRGBA[byte]()
		n := std.NormalizeRGBA32(c)
		o := std.ScaleToTypeRGBA32[byte](n)
		fmt.Println(c, n, o)
	}

	core.SetRandomNumberGenerator[byte](nil)

	for i := 0; i < 3; i++ {
		c := std.RandomRGBA[byte]()
		n := std.NormalizeRGBA32(c)
		o := std.ScaleToTypeRGBA32[byte](n)
		fmt.Println(c, n, o)
	}

	core.SetRandomNumberGenerator[byte](BadGenerator)

	for i := 0; i < 3; i++ {
		c := std.RandomRGBA[byte]()
		n := std.NormalizeRGBA32(c)
		o := std.ScaleToTypeRGBA32[byte](n)
		fmt.Println(c, n, o)
	}

	core.SetRandomNumberGenerator[byte](nil)

	for i := 0; i < 3; i++ {
		c := std.RandomRGBA[byte]()
		n := std.NormalizeRGBA32(c)
		o := std.ScaleToTypeRGBA32[byte](n)
		fmt.Println(c, n, o)
	}
}

func BadGenerator[T core.Numeric](r core.Tuple[T]) T {
	return 7
}
