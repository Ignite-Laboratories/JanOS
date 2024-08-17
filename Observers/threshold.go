package Observers

import (
	"github.com/ignite-laboratories/JanOS"
	"time"
)

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

func (o *ThresholdObserver) OnObservation(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	for i, pv := range ts.Data {
		if pv.Derivative > o.threshold {
			instant := ts.StartTime.Add(ts.Resolution.ToDuration(i))
			o.onTrigger(signal, instant, pv)
		}
	}
}
