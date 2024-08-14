package main

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"gonum.org/v1/gonum/mat"
	"log"
	"time"
)

var oscillator = Logic.NewOscillator(100, 1)
var sampler = Logic.NewSampler(32, OnSample, oscillator.Output, oscillator.Amplitude, oscillator.Frequency)

func main() {
	Logic.Drive(time.Nanosecond, oscillator, sampler)
}

func OnSample(dense mat.Dense) {
	log.Println("-----")
	log.Println(mat.Formatted(&dense))
}
