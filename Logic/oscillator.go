package Logic

import (
	"github.com/Ignite-Laboratories/JanOS/Logic/Math"
	"math"
	"time"
)

type Oscillator struct {
	LastUpdate   int64
	Output       *Dimension
	Amplitude    *Dimension
	Frequency    *Dimension
	phaseDegrees *Dimension
}

func NewOscillator(amplitude float64, frequency float64) *Oscillator {
	return &Oscillator{
		LastUpdate:   time.Now().UnixNano(),
		Output:       NewDimension("Output", Math.Y),
		Amplitude:    NewDimension("Amplitude", Math.Alpha, amplitude),
		Frequency:    NewDimension("Frequency", Math.Omega, frequency),
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

func (o *Oscillator) Cycle(periods int) []float64 {
	amplitude := o.Amplitude.Value
	frequency := o.Frequency.Value
	result := make([]float64, periods*360)

	var i = 0
	for period := 0; period < periods; period++ {
		for degree := 0; degree < 360; degree++ {
			phaseShiftInRadians := (float64(degree) * frequency) * math.Pi / 180
			angularFrequency := 2 * math.Pi * frequency
			value := amplitude * math.Sin(angularFrequency*time.Second.Seconds()+phaseShiftInRadians)
			result[i] = value
			i++
		}
	}
	return result
}
