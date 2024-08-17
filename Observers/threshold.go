package Observers

import (
	"github.com/ignite-laboratories/JanOS"
)

// ThresholdObserver samples a signal then calls the OnTrigger method whenever
// the derivative of the observed data crosses above a threshold.
type ThresholdObserver struct {
	Name      string
	threshold float64
	onTrigger func(observation JanOS.Observation)
}

// GetNamedValue returns the assigned name to this instance.
func (o *ThresholdObserver) GetNamedValue() string {
	return o.Name
}

// NewThresholdObserver samples a signal and calls the OnTrigger method whenever
// the derivative of the observed data crosses above a threshold.
func NewThresholdObserver(name string, threshold float64, onTrigger func(observation JanOS.Observation)) *ThresholdObserver {
	return &ThresholdObserver{
		Name:      name,
		threshold: threshold,
		onTrigger: onTrigger,
	}
}

func (o *ThresholdObserver) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	foundValues := make([]JanOS.InstantaneousValue, 0)
	for i, pv := range ts.Data {
		if pv.Derivative > o.threshold {
			instant := ts.StartTime.Add(ts.Resolution.ToDuration(i))
			foundValues = append(foundValues, JanOS.InstantaneousValue{
				Value:   pv,
				Instant: instant,
			})
		}
	}

	observation := JanOS.Observation{
		Observer: o,
		Signal:   signal,
		Values:   foundValues,
	}

	o.onTrigger(observation)
}
