package Observers

import "github.com/ignite-laboratories/JanOS"

// LoggingObserver is a simple "OnSample, print the info out" system for debugging.
type LoggingObserver struct{}

func (o *LoggingObserver) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	JanOS.Universe.Printf(signal, "Observation: %v", ts.Data)
}
