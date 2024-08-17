package JanOS

import (
	"github.com/ignite-laboratories/JanOS/Symbols"
	"time"
)

// Signal represents a mathematical value and its associated historical timeline.
type Signal struct {
	Name     string
	Symbol   Symbols.Symbol
	Timeline *timeline
}

// GetNamedValue returns the assigned name to this instance.
func (signal *Signal) GetNamedValue() string {
	return signal.Name
}

type signalManager struct {
	signals map[string]*Signal
}

func newSignalManager() *signalManager {
	return &signalManager{
		signals: make(map[string]*Signal),
	}
}

// GetNamedValue returns the assigned name to this instance.
func (mgr *signalManager) GetNamedValue() string {
	return "Signal"
}

// GetSignal references a previously stored signal.
func (mgr *signalManager) GetSignal(name string) *Signal {
	return mgr.signals[name]
}

// SetValue seeks to the appropriate position in time and replaces the values on the remainder of the buffer.
// The mentality is that when you set a value in time it will hold that value until it is changed.  We take
// time for granted in our environment, meaning that the setting of a value should always assume the future
// will hold that value ad infinitum.
func (signal *Signal) SetValue(instant time.Time, value float64) {
	Universe.Printf(Universe.Signals, "Set '%s' [%s] = %f", signal.Name, string(signal.Symbol), value)
	signal.Timeline.setValue(instant, value)
}

// GetValue seeks to the appropriate position in time and gets the value on the buffer at that instant.
func (signal *Signal) GetValue(instant time.Time) PointValue {
	return signal.Timeline.GetInstant(instant).Value
}

// NewSignal creates a new signal and sets its timeline to 0.
func (mgr *signalManager) NewSignal(name string, symbol Symbols.Symbol) *Signal {
	return mgr.NewSignalWithValue(name, symbol, 0)
}

// NewSignalWithValue creates a new signal and sets its timeline to the provided default value.
func (mgr *signalManager) NewSignalWithValue(name string, symbol Symbols.Symbol, defaultValue float64) *Signal {
	d := &Signal{
		Name:   name,
		Symbol: symbol,
	}
	d.Timeline = d.newTimeline(defaultValue)
	mgr.signals[name] = d
	Universe.Printf(d, "Let %s = %f | %dhz", string(symbol), defaultValue, d.Timeline.resolution.Frequency)
	return d
}
