package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/std/rgb"
	"github.com/ignite-laboratories/core/std/vector"
)

func main() {
	a1 := vector.From2D[int](1, 2)
	fmt.Println(a1)
	a2 := vector.From3D[int](1, 2, 3)
	fmt.Println(a2)
	a3 := vector.From4D[int](1, 2, 3, 4)
	fmt.Println(a3)
	a4 := vector.From5D[int](1, 2, 3, 4, 5)
	fmt.Println(a4)
	a5 := vector.From6D[int](1, 2, 3, 4, 5, 6)
	fmt.Println(a5)
	a6 := vector.From7D[int](1, 2, 3, 4, 5, 6, 7)
	fmt.Println(a6)

	c := rgb.FromHex[num.Morsel](0xAABBCCDD)
	fmt.Println(c)
}
