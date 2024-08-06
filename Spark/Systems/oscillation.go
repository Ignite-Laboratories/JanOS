package Systems

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/Spark"
	"github.com/Ignite-Laboratories/JanOS/Spark/Util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math"
	"time"
)

type OscillationSystem struct {
	Spark.Entity

	components OscillationSystemComponents
}

type OscillationSystemComponents struct {
	Oscillators *Spark.Components[Oscillator]
}

func NewOscillationSystem() OscillationSystem {
	return OscillationSystem{
		components: OscillationSystemComponents{
			Oscillators: &Spark.Components[Oscillator]{},
		},
	}
}

type Oscillator struct {
	Entity       Spark.Entity
	LastUpdate   int64
	Value        float64
	Amplitude    float64
	Frequency    float64
	Duration     time.Duration
	PhaseDegrees float64
}

func (sys OscillationSystem) StartOscillator(amplitude float64, frequency float64, duration time.Duration) Spark.Entity {
	oscillator := Oscillator{
		Entity:     Spark.Universe.CreateEntity(),
		LastUpdate: time.Now().UnixNano(),
		Amplitude:  amplitude,
		Frequency:  frequency,
		Duration:   duration,
	}
	sys.components.Oscillators.Set(oscillator.Entity, oscillator)
	return oscillator.Entity
}

func (sys OscillationSystem) GetOscillator(entity Spark.Entity) (Oscillator, bool) {
	return sys.components.Oscillators.Get(entity)
}

func (sys OscillationSystem) GetName() string         { return "Oscillation" }
func (sys OscillationSystem) GetEntity() Spark.Entity { return sys.Entity }

func (sys OscillationSystem) Initialize() {
}

func (sys OscillationSystem) Tick(inbox Spark.Inbox) {
	for _, oscillator := range sys.components.Oscillators.DB {
		now := time.Now().UnixNano()
		timeSinceLastStep := time.Duration(now - oscillator.LastUpdate)
		oscillator.PhaseDegrees += (timeSinceLastStep.Seconds() / oscillator.Duration.Seconds()) * 360
		if oscillator.PhaseDegrees > 360 {
			oscillator.PhaseDegrees -= 360
		}
		phaseShift := Util.DegreesToRadians(oscillator.PhaseDegrees)

		value := oscillator.Amplitude * math.Sin(((2*math.Pi)/oscillator.Frequency)*oscillator.Duration.Seconds()+phaseShift)
		oscillator.LastUpdate = now
		oscillator.Value = value
		sys.components.Oscillators.Set(oscillator.Entity, oscillator)
	}
}

func (sys OscillationSystem) OnDraw(entity Spark.Entity, screen *ebiten.Image) {
	oscillator, ok := sys.components.Oscillators.Get(entity)
	if ok {
		value := fmt.Sprintf("%f", oscillator.Value)
		ebitenutil.DebugPrintAt(screen, value, 5, int(oscillator.Entity)*10)
	}
}
