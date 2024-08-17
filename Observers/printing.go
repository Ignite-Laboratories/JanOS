package Observers

import "github.com/ignite-laboratories/JanOS"

// printingObserver is a simple "OnSample, print the info out" observer for debugging.
type printingObserver struct {
	Name string
}

func NewPrintingObserver(name string) *printingObserver {
	return &printingObserver{
		Name: name,
	}
}

// GetNamedValue returns the assigned name to this instance.
func (observer *printingObserver) GetNamedValue() string {
	return observer.Name
}

func (observer *printingObserver) OnSample(signal *JanOS.Signal, ts JanOS.TimeSlice) {
	JanOS.Universe.Printf(observer, "%s Observation: %v", signal.Symbol, ts.Data)
}
