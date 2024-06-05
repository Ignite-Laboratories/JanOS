package main

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/common"
	"github.com/Ignite-Laboratories/JanOS/common/config"
	"github.com/Ignite-Laboratories/JanOS/common/generate"
	"log"
)

func main() {
	config.Initialize()
	log.Println(`HEARTBEAT - NEUROLOGICAL SEED: ` + config.Current.Seed)

	ng := generate.NewNoiseGenerator(generate.NewNoiseType())
	go ng.Broadcast()
	pg := generate.NewPulseGenerator()
	go pg.Pulse()

	o := common.NewObserver(ng.Output, pg.Output)

	// Prints the observed data to stdout
	go func() {
		for line := range o.OutputStream {
			fmt.Println(line)
		}
	}()

	common.KeepAlive()
}
