package generate

import (
	"github.com/Ignite-Laboratories/JanOS/common/config"
	"github.com/google/uuid"
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

func NewPulseData(value string, delay time.Duration, duration time.Duration) PulseData {
	return PulseData{
		Value:    value,
		Delay:    delay,
		Duration: duration,
	}
}

func NewSeededPulseData() PulseData {
	return PulseData{
		Value:    uuid.New().String(),
		Delay:    time.Second,
		Duration: config.Current.DefaultDuration,
	}
}

func DefaultPulseData() PulseData {
	return PulseData{
		Value:    config.Current.Seed,
		Delay:    time.Second,
		Duration: config.Current.DefaultDuration,
	}
}

func (pg PulseGenerator) Pulse() {
	go func() {
		for {
			for i := 0; i < 50; i++ {
				pg.Output <- pg.Data.Value
				time.Sleep(pg.Data.Duration)
			}
			time.Sleep(pg.Data.Delay)
		}
	}()
}

func NewPulseGenerator(data ...PulseData) *PulseGenerator {
	pulseData := DefaultPulseData()
	if len(data) > 0 {
		pulseData = data[0]
	}

	return &PulseGenerator{
		Data:   pulseData,
		Output: make(chan string),
	}
}
