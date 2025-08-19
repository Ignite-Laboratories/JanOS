package main

import (
	"fmt"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/data"
)

type Comp struct {
	Value int
}

func main() {
	a := data.Contextual[int, std.Moment[int]]{}
}
