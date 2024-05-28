package main

import (
	"Heartbeat/config"
	"Heartbeat/generate"
	"log"
	"time"
)

func main() {
	config.Initialize()
	log.Println("SEED: " + config.Current.Seed)

	ng := generate.Noise(generate.NewNoiseType())
	pg := generate.Pulse(generate.NewPulseData())
	m := NewObserver()

	go m.Observe(ng.Output)
	go m.Observe(pg.Output)

	//NewNeuron(m)
	//
	for {
		time.Sleep(time.Second)
	}
}
