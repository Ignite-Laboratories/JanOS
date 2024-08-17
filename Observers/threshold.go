package Observers

import (
	"github.com/ignite-laboratories/JanOS"
	"time"
)

type thresholdObserver struct {
	threshold float64
	onTrigger func(signal *JanOS.Signal, instant time.Time)
}

func NewThresholdObserver(threshold float64, onTrigger func(signal *JanOS.Signal, instant time.Time)) *thresholdObserver {
	return &thresholdObserver{
		threshold: threshold,
		onTrigger: onTrigger,
	}
}

func (o *thresholdObserver) OnObservation(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	for i, pv := range ts.Data {
		if pv.Derivative > o.threshold {
			o.onTrigger(signal, ts.StartTime.Add(ts.Resolution.ToDuration(i)))
		}
	}
}
