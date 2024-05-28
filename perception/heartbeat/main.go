package main

import (
	"github.com/Project-Arwen/Rivendell/perception/heartbeat/config"
	"github.com/Project-Arwen/Rivendell/perception/heartbeat/generate"
	"log"
	"time"
)

type Neuron struct {
	Seed   string
	Stream chan string
}

func NewNeuron(m *Observer) *Neuron {
	n := Neuron{
		Seed:   "!",
		Stream: make(chan string),
	}
	go n.Observe(m.Stream)

	return &n
}

func (n Neuron) Observe(source chan string) {
	for msg := range source {
		if msg == n.Seed {
			log.Println("PULSE IDENTIFIED")
		}
	}
}

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
