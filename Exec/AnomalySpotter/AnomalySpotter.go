package main

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"log"
)

var oscillator = Logic.NewOscillator(100, 1)

func main() {
	result := oscillator.Cycle(1)
	log.Println(result)
}
