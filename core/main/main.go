package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/sys/pad"
	"github.com/ignite-laboratories/core/sys/pad/scheme"
	"math"
)

func main() {
	fmt.Println(pad.String[orthogonal.Left, scheme.Reverse](10, "11111", "ABC"))
	fmt.Println(pad.String[orthogonal.Left, scheme.Tile](10, "11111", "ABC"))
	fmt.Println(pad.String[orthogonal.Right, scheme.Reverse](10, "11111", "ABC"))
	fmt.Println(pad.String[orthogonal.Right, scheme.Tile](10, "11111", "ABC"))

	math.min
}
