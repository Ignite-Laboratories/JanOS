package main

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/config"
	"github.com/Ignite-Laboratories/JanOS/generate"
	"github.com/Ignite-Laboratories/JanOS/support"
	"log"
	"time"
)

func main() {
	config.Initialize()
	log.Println(`HEARTBEAT - NEUROLOGICAL SEED: ` + config.Current.Seed)

	o := support.NewObserver()

	// Prints the observed data to stdout
	go func() {
		for line := range o.OutputStream {
			fmt.Println(line)
		}
	}()

	ng := generate.Noise(generate.NewNoiseType())
	ng.Broadcast()
	go o.MuxChannel(ng.Output)

	pg := generate.NewPulseGenerator()
	pg.Pulse()
	go o.MuxChannel(pg.Output)

	for {
		time.Sleep(time.Second)
	}
}
