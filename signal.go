package JanOS

import (
	"github.com/ignite-laboratories/JanOS/Symbols"
	"math"
	"time"
)

// Signal represents a mathematical value and its associated historical timeline.
type Signal struct {
	Name     string
	Symbol   Symbols.Symbol
	Timeline *timeline
	OnStep   func(signal *Signal, iv InstantaneousValue) float64
}

// GetNamedValue returns the assigned name to this instance.
func (signal *Signal) GetName() string {
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
func (mgr *signalManager) GetName() string {
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

// GetInstantValue seeks to the appropriate position in time and gets the value on the buffer at that instant.
func (signal *Signal) GetInstantValue(instant time.Time) InstantaneousValue {
	return signal.Timeline.GetInstant(instant)
}

// NewToggleSignal creates a signal that toggles between a positive and negative value at the provided frequency.
// You can adjust its value to directly control how frequently the system "toggles" between "high" and "low".
func (mgr *signalManager) NewToggleSignal(name string, symbol Symbols.Symbol, frequency int) *Signal {
	lastPeriod := time.Now()

	outputSignal := mgr.NewSignalWithValue(name, symbol, float64(frequency), func(signal *Signal, iv InstantaneousValue) float64 {
		// Recalculate the resolution in case the signal changed
		newFrequency := int(math.Abs(iv.Point.Value))
		resolution := NewResolution(newFrequency)

		// Return the inverse if it is time
		if iv.Instant.Sub(lastPeriod) > resolution.Duration {
			lastPeriod = iv.Instant
			return iv.Point.Value * -1
		}

		// Return no change if it's not time
		return iv.Point.Value
	})

	Universe.Printf(outputSignal, "Let ƒ(%s) => -%s | %v", string(symbol), string(symbol), NewResolution(frequency).Duration)

	return outputSignal
}

// Mux creates a signal that uses the provided formula to mux the provided signals together.
func (mgr *signalManager) Mux(name string, symbol Symbols.Symbol, formula Formula, signals ...*Signal) *Signal {
	outputSignal := Universe.Signals.NewSignal(name, symbol, func(signal *Signal, iv InstantaneousValue) float64 {
		sourceValue := 0.0
		otherValues := make([]float64, len(signals))
		for i, s := range signals {
			otherValues[i] = s.GetInstantValue(iv.Instant).Point.Value
		}

		return formula.Operation(sourceValue, otherValues...)
	})

	symbols := SpacedStringSet(formula.Operator, Select(signals, func(s *Signal) string { return string(s.Symbol) })...)
	Universe.Printf(outputSignal, "Let ƒ(%s) => y = %s", string(symbol), symbols)
	return outputSignal
}

// NewSignal creates a new signal and sets its timeline to 0.
func (mgr *signalManager) NewSignal(name string, symbol Symbols.Symbol, onStep func(*Signal, InstantaneousValue) float64) *Signal {
	return mgr.NewSignalWithValue(name, symbol, 0, onStep)
}

// NewSignalWithValue creates a new signal and sets its timeline to the provided default value.
func (mgr *signalManager) NewSignalWithValue(name string, symbol Symbols.Symbol, defaultValue float64, onStep func(*Signal, InstantaneousValue) float64) *Signal {
	s := &Signal{
		Name:   name,
		Symbol: symbol,
		OnStep: onStep,
	}
	s.Timeline = s.newTimeline(defaultValue, onStep)
	mgr.signals[name] = s
	if defaultValue != 0 {
		Universe.Printf(s, "Let %s = %f | %dhz", string(symbol), defaultValue, s.Timeline.resolution.Frequency)
	} else {
		Universe.Printf(s, "Let %s exist | %dhz", string(symbol), s.Timeline.resolution.Frequency)
	}
	return s
}
