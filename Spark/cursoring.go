package Spark

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"time"
)

type CursoringSystem struct {
	Entity

	components CursoringSystemComponents
}

type CursoringSystemComponents struct {
	Cursors *Components[Cursor]
}

func NewCursoringSystem() CursoringSystem {
	return CursoringSystem{
		components: CursoringSystemComponents{
			Cursors: &Components[Cursor]{},
		},
	}
}

type Cursor struct {
	Entity           Entity
	OscillatorEntity Entity
	LastUpdate       time.Time
	Buffer           []float64
	BufferSize       int64
	Duration         time.Duration
	interval         time.Duration
}

func (sys CursoringSystem) GetCursorBuffer(entity Entity) []float64 {
	c, _ := sys.components.Cursors.Get(entity)
	return c.Buffer
}

func (sys CursoringSystem) StartCursor(oscillator Entity, bufferSize int64, duration time.Duration) Entity {
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

func (sys CursoringSystem) GetName() string   { return "Cursoring" }
func (sys CursoringSystem) GetEntity() Entity { return sys.Entity }

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

func (sys CursoringSystem) OnDraw(entity Entity, screen *ebiten.Image) {
	cursor, ok := sys.components.Cursors.Get(entity)
	if ok {
		value := fmt.Sprintf("%f", cursor.Buffer)
		ebitenutil.DebugPrintAt(screen, value, 5, int(cursor.Entity)*10)
	}
}
