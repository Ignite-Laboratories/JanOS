package Observers

import "github.com/ignite-laboratories/JanOS"

// PrintingObserver is a simple "OnSample, print the info out" observer for debugging.
type PrintingObserver struct{}

func (o *PrintingObserver) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	JanOS.Universe.Printf(signal, "Observation: %v", ts.Data)
}
