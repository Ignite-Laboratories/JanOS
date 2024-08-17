package Observers

import (
	"github.com/ignite-laboratories/JanOS"
	"time"
)

// IntegralObserver calculates the integral of each provided TimeSlice on observation.
type IntegralObserver struct {
	onTrigger func(signal *JanOS.Signal, instant time.Time, value float64)
}

// NewIntegralObserver calculates the integral of each provided TimeSlice on observation.
func NewIntegralObserver(onTrigger func(signal *JanOS.Signal, instant time.Time, value float64)) *IntegralObserver {
	return &IntegralObserver{
		onTrigger: onTrigger,
	}
}

func (o *IntegralObserver) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	o.onTrigger(signal, ts.StartTime, ts.Integrate())
}
