package Observers

import (
	"github.com/ignite-laboratories/JanOS"
	"time"
)

// ThresholdObserver samples data and then calls the OnTrigger method whenever
// the derivative of the observed data crosses above a threshold.
type ThresholdObserver struct {
	threshold float64
	onTrigger func(signal *JanOS.Signal, instant time.Time, pointValue JanOS.PointValue)
}

func NewThresholdObserver(threshold float64, onTrigger func(signal *JanOS.Signal, instant time.Time, pointValue JanOS.PointValue)) *ThresholdObserver {
	return &ThresholdObserver{
		threshold: threshold,
		onTrigger: onTrigger,
	}
}

func (o *ThresholdObserver) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	for i, pv := range ts.Data {
		if pv.Derivative > o.threshold {
			instant := ts.StartTime.Add(ts.Resolution.ToDuration(i))
			o.onTrigger(signal, instant, pv)
		}
	}
}
