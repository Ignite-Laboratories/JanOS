package JanOS

import (
	"github.com/ignite-laboratories/JanOS/util"
	"log"
	"math"
	"sync"
	"time"
)

type Signal struct {
	util.Named
	Name       string
	Symbol     string
	Value      float64
	Derivative float64
	function   func(value float64, derivative float64) float64

	lock        sync.Mutex
	newValue    float64
	hasNewValue bool
}

func (signal *Signal) SetValue(value float64) *Signal {
	signal.newValue = value
	signal.hasNewValue = true
	return signal
}

// onStep is called on every resolution step by the main timekeeper
func (signal *Signal) onStep() {
	signal.lock.Lock()
	defer signal.lock.Unlock()

	oldValue := signal.Value

	// Check if we have a new value and set it
	if signal.hasNewValue {
		signal.Derivative = signal.newValue - oldValue
		signal.Value = signal.newValue
		signal.hasNewValue = false
	}

	// If this signal has a function, call it
	if signal.function != nil {
		newValue := signal.function(signal.Value, signal.Derivative)
		signal.Derivative = newValue - oldValue
		signal.Value = newValue
	}
}

// WithFunction sets the function to perform on this signal whenever a time step occurs.
// This is a destructive operation as signals are meant to only have one function or
// none - if you would like to perform more operations create a new function that
// performs the steps you desire.
func (signal *Signal) WithFunction(function func(value float64, derivative float64) float64) *Signal {
	signal.function = function
	return signal
}

// CreateMux creates a new signal that uses the provided formula to mux the provided signals together.
func (signal *Signal) CreateMux(formula util.Formula, signals ...*Signal) *Signal {
	outputSignal := Signals.Create().WithFunction(func(value float64, derivative float64) float64 {
		sourceValue := signal.Value
		otherValues := make([]float64, len(signals))
		for i, s := range signals {
			otherValues[i] = s.Value
		}

		return formula.Operation(sourceValue, otherValues...)
	})

	//symbols := util.SpacedStringSet(formula.Operator, util.Select(signals, func(s *Signal) string { return s.Symbol })...)
	//Logging.Printf(outputSignal, "ƒ(%s) => y = %s", symbol, symbols)
	return outputSignal
}

// Toggle causes the signal to toggle between a positive and negative value at the provided frequency.
// The periodicity of this is determined by the inherent value of the signal, in Hz.
func (signal *Signal) Toggle() *Signal {
	lastPeriod := time.Now()

	return signal.WithFunction(func(value float64, derivative float64) float64 {
		now := time.Now()

		// We calculate the periodicity as a function of the current value on the signal
		nanoSecs := int64((1 / math.Abs(value)) * float64(time.Second))
		period := time.Duration(nanoSecs)

		// If we are above the currently calculated period...
		if now.Sub(lastPeriod) > period {
			lastPeriod = now
			// ...flip the sign
			return -value
		}

		// ...otherwise, send the same value back
		return value
	})
}

// AbsoluteValue creates a new signal that maintains the absolute value of the source signal.
func (signal *Signal) AbsoluteValue() *Signal {
	lastTime := time.Now()

	outputSignal := Signals.Create().WithFunction(func(value float64, derivative float64) float64 {
		if signal.Derivative != 0 {
			log.Println(time.Since(lastTime))
			lastTime = time.Now()
		}

		return math.Abs(signal.Value)
	})

	//	Logging.Printf(outputSignal, "ƒ(%s) => y = |y| (Absolute Value)", outputSignal.Symbol)
	return outputSignal
}
