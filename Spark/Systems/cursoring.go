package Systems

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/Spark"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	LastUpdate       time.Time
	Buffer           []float64
	BufferSize       int64
	Duration         time.Duration
	interval         time.Duration
}

func (sys CursoringSystem) GetCursorBuffer(entity Spark.Entity) []float64 {
	c, _ := sys.components.Cursors.Get(entity)
	return c.Buffer
}

func (sys CursoringSystem) StartCursor(oscillator Spark.Entity, bufferSize int64, duration time.Duration) Spark.Entity {
	cursor := Cursor{
		Entity:           Spark.Universe.CreateEntity(),
		OscillatorEntity: oscillator,
		LastUpdate:       time.Now(),
		Buffer:           make([]float64, bufferSize),
		BufferSize:       bufferSize,
		Duration:         duration,
		interval:         time.Duration(duration.Nanoseconds() / bufferSize),
	}
	sys.components.Cursors.Set(cursor.Entity, cursor)
	return cursor.Entity
}

func (sys CursoringSystem) GetName() string         { return "Cursoring" }
func (sys CursoringSystem) GetEntity() Spark.Entity { return sys.Entity }

func (sys CursoringSystem) Initialize() {
}

func (sys CursoringSystem) Tick(inbox Spark.Inbox) {
	for _, cursor := range sys.components.Cursors.DB {
		if oscillationSystem, ok := Spark.Universe.GetSystem(OscillationSystem{}).(OscillationSystem); ok {
			o, _ := oscillationSystem.GetOscillator(cursor.OscillatorEntity)

			now := time.Now()
			timeSinceLastUpdate := now.Sub(cursor.LastUpdate)

			if timeSinceLastUpdate.Nanoseconds() > cursor.interval.Nanoseconds() {
				cursor.Buffer = append(cursor.Buffer[1:], o.Value)
				cursor.LastUpdate = now

				sys.components.Cursors.Set(cursor.Entity, cursor)
			}
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
