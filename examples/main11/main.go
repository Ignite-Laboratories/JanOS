package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/sys/atlas"
	"time"
)

func main() {
	defer core.Exit()

	core.Verbose = true
	go core.Shutdown(time.Second * 3)

	atlas.Impulse.Spark()
}
