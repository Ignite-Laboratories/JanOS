package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/std/rgba"
)

func main() {
	a := rgba.FromHex[byte](0xaabbccdd)
	b := rgba.FromHex[num.Nibble](0xaabbccdd)
	fmt.Println(a)
	fmt.Println(b)
}
