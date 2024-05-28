package generate

import (
	"github.com/Project-Arwen/Rivendell/perception/heartbeat/config"
	"time"
)

type PulseGenerator struct {
	Data   PulseData
	Output chan string
}

type PulseData struct {
	Value    string
	Duration time.Duration
	Delay    time.Duration
}

func NewPulseData() PulseData {
	return PulseData{
		Value:    config.Current.Seed,
		Delay:    time.Second,
		Duration: config.Current.DefaultDuration,
	}
}

func (pg PulseGenerator) Pulse() {
	for {
		for i := 0; i < 5; i++ {
			pg.Output <- pg.Data.Value
			time.Sleep(pg.Data.Duration)
		}
		time.Sleep(pg.Data.Delay)
	}
}

func Pulse(data PulseData) *PulseGenerator {
	pg := PulseGenerator{
		Data:   data,
		Output: make(chan string),
	}
	go pg.Pulse()
	return &pg
}
