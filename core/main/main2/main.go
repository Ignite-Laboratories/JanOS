package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/std/rgb"
)

func main() {
	a := rgb.FromHex[byte](0xaabbcc00)
	b := rgb.FromHex[num.Nibble](0xaabbcc00)
	fmt.Println(a)
	fmt.Println(b)
}
