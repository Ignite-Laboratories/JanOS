package JanOS

import (
	"math"
	"time"
)

type signalManager struct {
	signals []*Signal
}

func (mgr *signalManager) GetName() string { return "Signals" }

func newSignalManager() *signalManager {
	mgr := &signalManager{
		signals: make([]*Signal, 0),
	}

	lastStep := time.Now()

	go func() {
		for {
			if Terminate {
				break
			}

			now := time.Now()

			if now.Sub(lastStep) > StepSize() {
				lastStep = now

				for _, signal := range mgr.signals {
					signal.onStep()
				}
			}
			time.Sleep(5)
		}
		Logging.Printf(mgr, "Signal manager shutdown")
	}()

	return mgr
}

func (mgr *signalManager) Create() *Signal {
	signal := &Signal{}
	mgr.signals = append(mgr.signals, signal)

	//	Logging.Printf(signal, "Let %s exist | %dhz", symbol, Frequency)
	return signal
}

func (mgr *signalManager) SineWave(amplitude *Signal, frequency *Signal) *Signal {
	lastPeriod := time.Now()

	outputSignal := mgr.Create().WithFunction(func(value float64, derivative float64) float64 {
		f := frequency.Value
		a := amplitude.Value
		// Seconds() gives us a float, which acts as a scale factor
		// for the position of the period relative to 1 second.
		periodOffset := time.Since(lastPeriod).Seconds()

		phaseShiftInRadians := (360.0 * periodOffset * f) * math.Pi / 180
		angularFrequency := 2 * math.Pi * f
		calculatedValue := a * math.Sin(angularFrequency+phaseShiftInRadians)
		return calculatedValue
	})
	//	Logging.Printf(outputSignal, "ƒ(%s) => y = %s * sin(%s * t) (Sine Wave)", outputSignal.Symbol, amplitude.Symbol, frequency.Symbol)

	return outputSignal
}
