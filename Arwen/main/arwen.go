package main

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/Arwen"
	"gonum.org/v1/gonum/mat"
	"log"
	"time"
)

var oscillator = Arwen.NewStdOscillator(100, 1)
var sampler = Arwen.NewSampler(32, OnSample, oscillator.Output, oscillator.Amplitude, oscillator.Frequency)

func main() {
	go print()
	go adjust()

	for {

	}
}

func print() {
	for {
		go oscillator.Tick()
		go sampler.Tick()
		time.Sleep(time.Nanosecond)
	}
}

func OnSample(dense mat.Dense) {
	log.Println("-----")
	log.Println(mat.Formatted(&dense))
}

func printDimension(d *Arwen.Dimension) string {
	return fmt.Sprintf("%s - %s - %f", d.Name, d.Symbol, d.Value)
}

func adjust() {
	for {
		time.Sleep(time.Second * 5)
		oscillator.Frequency.Value = oscillator.Frequency.Value * 1.25
	}
}
