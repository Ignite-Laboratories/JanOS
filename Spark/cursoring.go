package Spark

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"time"
)

type CursoringSystem struct {
	Logic.Entity

	components CursoringSystemComponents
}

type CursoringSystemComponents struct {
	Cursors *Logic.Components[Cursor]
}

func NewCursoringSystem() CursoringSystem {
	return CursoringSystem{
		components: CursoringSystemComponents{
			Cursors: &Logic.Components[Cursor]{},
		},
	}
}

type Cursor struct {
	Entity           Logic.Entity
	OscillatorEntity Logic.Entity
	LastUpdate       time.Time
	Buffer           []float64
	BufferSize       int64
	Duration         time.Duration
	interval         time.Duration
}

func (sys CursoringSystem) GetCursorBuffer(entity Logic.Entity) []float64 {
	c, _ := sys.components.Cursors.Get(entity)
	return c.Buffer
}

func (sys CursoringSystem) StartCursor(oscillator Logic.Entity, bufferSize int64, duration time.Duration) Logic.Entity {
	cursor := Cursor{
		Entity:           Universe.CreateEntity(),
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
func (sys CursoringSystem) GetEntity() Logic.Entity { return sys.Entity }

func (sys CursoringSystem) Initialize() {
}

func (sys CursoringSystem) Tick(inbox Inbox) {
	for _, cursor := range sys.components.Cursors.DB {
		o, _ := Universe.Oscillation.GetOscillator(cursor.OscillatorEntity)

		now := time.Now()
		timeSinceLastUpdate := now.Sub(cursor.LastUpdate)

		if timeSinceLastUpdate.Nanoseconds() > cursor.interval.Nanoseconds() {
			cursor.Buffer = append(cursor.Buffer[1:], o.Value)
			cursor.LastUpdate = now

			sys.components.Cursors.Set(cursor.Entity, cursor)
		}
	}
}

func (sys CursoringSystem) OnDraw(entity Logic.Entity, screen *ebiten.Image) {
	cursor, ok := sys.components.Cursors.Get(entity)
	if ok {
		value := fmt.Sprintf("%f", cursor.Buffer)
		ebitenutil.DebugPrintAt(screen, value, 5, int(cursor.Entity)*10)
	}
}
