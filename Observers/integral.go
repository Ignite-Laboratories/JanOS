package Observers

import (
	"github.com/ignite-laboratories/JanOS"
	"time"
)

// integralObserver calculates the integral of each provided TimeSlice on observation.
type integralObserver struct {
	Name      string
	onTrigger func(signal *JanOS.Signal, instant time.Time, value float64)
}

// GetNamedValue returns the assigned name to this instance.
func (observer *integralObserver) GetNamedValue() string {
	return observer.Name
}

// NewIntegralObserver calculates the integral of each provided TimeSlice on observation.
func NewIntegralObserver(name string, onTrigger func(signal *JanOS.Signal, instant time.Time, value float64)) *integralObserver {
	return &integralObserver{
		Name:      name,
		onTrigger: onTrigger,
	}
}

func (observer *integralObserver) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	observer.onTrigger(signal, ts.StartTime, ts.Integrate())
}
