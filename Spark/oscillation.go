package Spark

import (
	"github.com/Ignite-Laboratories/JanOS/Spark/Util"
	"math"
	"time"
)

type OscillationSystem struct {
	Entity

	components OscillationSystemComponents
}

type OscillationSystemComponents struct {
	Oscillators *Components[Oscillator]
}

func NewOscillationSystem() OscillationSystem {
	return OscillationSystem{
		components: OscillationSystemComponents{
			Oscillators: &Components[Oscillator]{},
		},
	}
}

type Oscillator struct {
	Entity       Entity
	LastUpdate   int64
	Value        float64
	Amplitude    float64
	Frequency    float64
	Duration     time.Duration
	PhaseDegrees float64
}

func (sys OscillationSystem) StartOscillator(amplitude float64, frequency float64, duration time.Duration) Entity {
	oscillator := Oscillator{
		Entity:     Universe.CreateEntity(),
		LastUpdate: time.Now().UnixNano(),
		Amplitude:  amplitude,
		Frequency:  frequency,
		Duration:   duration,
	}
	sys.components.Oscillators.Set(oscillator.Entity, oscillator)
	return oscillator.Entity
}

func (sys OscillationSystem) GetOscillator(entity Entity) (Oscillator, bool) {
	return sys.components.Oscillators.Get(entity)
}

func (sys OscillationSystem) GetName() string   { return "Oscillation" }
func (sys OscillationSystem) GetEntity() Entity { return sys.Entity }

func (sys OscillationSystem) Initialize() {
}

func (sys OscillationSystem) Tick(inbox Inbox) {
	for _, oscillator := range sys.components.Oscillators.DB {
		now := time.Now().UnixNano()
		timeSinceLastStep := time.Duration(now - oscillator.LastUpdate)
		oscillator.PhaseDegrees += (timeSinceLastStep.Seconds() / oscillator.Duration.Seconds()) * 360
		if oscillator.PhaseDegrees > 360 {
			oscillator.PhaseDegrees -= 360
		}
		phaseShift := Util.DegreesToRadians(oscillator.PhaseDegrees * oscillator.Frequency)

		value := oscillator.Amplitude * math.Sin(((2*math.Pi)*oscillator.Frequency)*oscillator.Duration.Seconds()+phaseShift)
		oscillator.LastUpdate = now
		oscillator.Value = value
		sys.components.Oscillators.Set(oscillator.Entity, oscillator)
	}
}

//func (sys OscillationSystem) OnDraw(entity Spark.Entity, screen *ebiten.Image) {
//	oscillator, ok := sys.components.Oscillators.Get(entity)
//	if ok {
//		value := fmt.Sprintf("%f", oscillator.Value)
//		ebitenutil.DebugPrintAt(screen, value, 5, int(oscillator.Entity)*10)
//	}
//}
