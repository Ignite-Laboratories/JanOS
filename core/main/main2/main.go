package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/bounded"
)

func main() {
	a := bounded.By[uint64](8, 2, 15)
	fmt.Println(a)

	n := a.Normalize()
	fmt.Println(n)
	n32 := a.Normalize32()
	fmt.Println(n32)

	a.SetFromNormalized(n)
	fmt.Println(a)

	a.SetFromNormalized32(n32)
	fmt.Println(a)
}
