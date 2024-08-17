package Observers

import (
	"github.com/ignite-laboratories/JanOS"
)

// ThresholdObserver samples a signal then calls the OnTrigger method whenever
// the derivative of the observed data crosses above a threshold.
type ThresholdObserver struct {
	threshold float64
	onTrigger func(signal *JanOS.Signal, foundValues []JanOS.InstantaneousValue)
}

// NewThresholdObserver samples a signal and calls the OnTrigger method whenever
// the derivative of the observed data crosses above a threshold.
func NewThresholdObserver(threshold float64, onTrigger func(signal *JanOS.Signal, foundValues []JanOS.InstantaneousValue)) *ThresholdObserver {
	return &ThresholdObserver{
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

	o.onTrigger(signal, foundValues)
}
