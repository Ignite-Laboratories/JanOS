package main

import (
	"fmt"
	"time"

	"git.ignitelabs.net/janos/core"
	"git.ignitelabs.net/janos/core/sys/atlas"
)

func main() {
	core.WhileAlive(func() {
		fmt.Println(atlas.Verbose())
	}, time.Second)
}

type rawr struct {
	Precision uint
}
