package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

func main() {
	for {
		b, _ := bounded.By[num.Nibble](0, 0, 5)
		a := std.NewUniqueSet(b)

		for {
			result := a.Random(15)
			fmt.Println(result)
		}

	}
}
