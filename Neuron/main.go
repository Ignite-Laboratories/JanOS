package main

import (
	"Neuron/config"
	"log"
	"time"
)

func main() {
	config.Initialize()
	log.Println(`NEURON - NEUROLOGICAL SEED: ` + config.Current.Seed)

	// This will read all stdin lines and put them on a channel
	o := NewStdInObserver()

	go func() {
		for line := range o.OutputStream {
			log.Println("NEURON: " + line)
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}
