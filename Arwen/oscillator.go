package Arwen

import (
	"github.com/Ignite-Laboratories/JanOS/Logic/Math"
	"math"
	"time"
)

func NewStdOscillator(amplitude float64, frequency float64) *Oscillator {
	return NewOscillator(
		NewDimension("Output", Math.Y),
		NewDimension("Amplitude", Math.Alpha, amplitude),
		NewDimension("Frequency", Math.Omega, frequency),
	)
}

type Oscillator struct {
	LastUpdate   int64
	Output       *Dimension
	Amplitude    *Dimension
	Frequency    *Dimension
	phaseDegrees *Dimension
}

func NewOscillator(output *Dimension, amplitude *Dimension, frequency *Dimension) *Oscillator {
	return &Oscillator{
		LastUpdate:   time.Now().UnixNano(),
		Output:       output,
		Amplitude:    amplitude,
		Frequency:    frequency,
		phaseDegrees: NewDimension("Phase Degrees", Math.Phi),
	}
}

func (o *Oscillator) Tick() {
	amplitude := o.Amplitude.Value
	frequency := o.Frequency.Value
	phaseDegrees := o.phaseDegrees.Value

	now := time.Now().UnixNano()
	elapsedTime := time.Duration(now - o.LastUpdate).Seconds()

	phaseDegrees += (elapsedTime / time.Second.Seconds()) * 360
	phaseShiftInRadians := (phaseDegrees * frequency) * math.Pi / 180
	angularFrequency := 2 * math.Pi * frequency
	value := amplitude * math.Sin(angularFrequency*time.Second.Seconds()+phaseShiftInRadians)

	o.Output.Value = value
	o.phaseDegrees.Value = phaseDegrees
	o.LastUpdate = now
}
