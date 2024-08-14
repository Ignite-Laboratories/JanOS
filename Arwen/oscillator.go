package Arwen

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/Ignite-Laboratories/JanOS/Logic/Math"
	"math"
	"time"
)

type Oscillator struct {
	LastUpdate   int64
	Output       *Logic.Dimension
	Amplitude    *Logic.Dimension
	Frequency    *Logic.Dimension
	phaseDegrees *Logic.Dimension
}

func NewOscillator(amplitude float64, frequency float64) *Oscillator {
	return &Oscillator{
		LastUpdate:   time.Now().UnixNano(),
		Output:       Logic.NewDimension("Output", Math.Y),
		Amplitude:    Logic.NewDimension("Amplitude", Math.Alpha, amplitude),
		Frequency:    Logic.NewDimension("Frequency", Math.Omega, frequency),
		phaseDegrees: Logic.NewDimension("Phase Degrees", Math.Phi),
	}
}

func (o *Oscillator) Calculate(resolution int) []float64 {
	amplitude := o.Amplitude.Value
	frequency := o.Frequency.Value
	stepDegrees := 360 / resolution
	result := make([]float64, resolution)

	for i := 0; i < resolution; i++ {
		phaseShiftInRadians := (float64(stepDegrees*i) * frequency) * math.Pi / 180
		angularFrequency := 2 * math.Pi * frequency
		value := amplitude * math.Sin(angularFrequency*time.Second.Seconds()+phaseShiftInRadians)
		result[i] = value
		i++
	}

	return result
}
