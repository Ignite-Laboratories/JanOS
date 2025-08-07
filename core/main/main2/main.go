package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/bounded"
	"github.com/ignite-laboratories/core/std/num"
)

func main() {
	a := bounded.By[num.Morsel](22, 33, 66, true)
	fmt.Println(a)
}
