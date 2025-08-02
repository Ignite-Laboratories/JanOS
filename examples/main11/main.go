package main

import (
	"github.com/ignite-laboratories/core"
	"time"
)

func main() {
	defer core.Exit()

	core.Verbose = true
	go core.Shutdown(time.Second * 3)

	core.Impulse.Spark()
}
