package Systems

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/Spark"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"time"
)

type CursoringSystem struct {
	Spark.Entity

	components CursoringSystemComponents
}

type CursoringSystemComponents struct {
	Cursors *Spark.Components[Cursor]
}

func NewCursoringSystem() CursoringSystem {
	return CursoringSystem{
		components: CursoringSystemComponents{
			Cursors: &Spark.Components[Cursor]{},
		},
	}
}

type Cursor struct {
	Entity           Spark.Entity
	OscillatorEntity Spark.Entity
	LastUpdate       int64
	Buffer           []any
	Index            int64
	Resolution       int64
	DutyCycle        time.Duration
}

func (sys CursoringSystem) StartCursor(oscillator Spark.Entity, resolution int64, dutyCycle time.Duration) {
	cursor := Cursor{
		Entity:           Spark.Universe.CreateEntity(),
		OscillatorEntity: oscillator,
		LastUpdate:       time.Now().UnixNano(),
		Buffer:           make([]any, resolution),
		Index:            0,
		Resolution:       resolution,
		DutyCycle:        dutyCycle,
	}
	sys.components.Cursors.Set(cursor.Entity, cursor)
}

func (sys CursoringSystem) GetName() string         { return "Cursoring" }
func (sys CursoringSystem) GetEntity() Spark.Entity { return sys.Entity }

func (sys CursoringSystem) Initialize() {
}

func (sys CursoringSystem) Tick(inbox Spark.Inbox) {
	for _, cursor := range sys.components.Cursors.DB {
		if oscillationSystem, ok := Spark.Universe.GetSystem(OscillationSystem{}).(OscillationSystem); ok {
			o, _ := oscillationSystem.GetOscillator(cursor.OscillatorEntity)
			now := time.Now().UnixNano()
			timeSinceLastStep := time.Duration(now - cursor.LastUpdate).Nanoseconds()
			resolutionInterval := cursor.DutyCycle.Nanoseconds() / cursor.Resolution

			if timeSinceLastStep > resolutionInterval {
				cursor.Buffer[cursor.Index] = o.Value
				cursor.Index++
			}
			if cursor.Index >= cursor.Resolution {
				cursor.Index = 0
				log.Println(fmt.Sprint(cursor.Buffer))
			}
			cursor.LastUpdate = now
			sys.components.Cursors.Set(cursor.Entity, cursor)
		}
	}
}

func (sys CursoringSystem) OnDraw(entity Spark.Entity, screen *ebiten.Image) {
	cursor, ok := sys.components.Cursors.Get(entity)
	if ok {
		value := fmt.Sprintf("%f", cursor.Buffer)
		ebitenutil.DebugPrintAt(screen, value, 5, int(cursor.Entity)*10)
	}
}
