package generate

import (
	"Heartbeat/config"
	support "Support"
	"time"
)

type NoiseGenerator struct {
	Output chan string
	Type   NoiseType
}

type NoiseType struct {
	Width    int
	Duration time.Duration
}

func (nt *NoiseType) CreateNoise() string {
	return support.RandomString(nt.Width)
}

func (ng NoiseGenerator) Broadcast() {
	for {
		ng.Output <- ng.Type.CreateNoise()
		time.Sleep(ng.Type.Duration)
	}
}

func NewNoiseType() NoiseType {
	return NoiseType{
		Width:    36,
		Duration: config.Current.DefaultDuration,
	}
}

func Noise(nt NoiseType) *NoiseGenerator {
	ng := NoiseGenerator{
		Output: make(chan string),
		Type:   nt,
	}
	go ng.Broadcast()
	return &ng
}
