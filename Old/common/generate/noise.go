package generate

import (
	"github.com/Ignite-Laboratories/JanOS/common"
	"github.com/Ignite-Laboratories/JanOS/common/config"
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

func (ng NoiseGenerator) Broadcast() {
	for {
		ng.Output <- common.RandomString(ng.Type.Width)
		time.Sleep(ng.Type.Duration)
	}
}

func NewNoiseType() NoiseType {
	return NoiseType{
		Width:    36,
		Duration: config.Current.DefaultDuration,
	}
}

func NewNoiseGenerator(nt NoiseType) *NoiseGenerator {
	ng := NoiseGenerator{
		Output: make(chan string),
		Type:   nt,
	}
	go ng.Broadcast()
	return &ng
}
