package Observers

import "github.com/ignite-laboratories/JanOS"

type LoggingObserver struct{}

func (o *LoggingObserver) OnObservation(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	JanOS.Universe.Printf(signal, "Observation: %v", ts.Data)
}
