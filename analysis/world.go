package analysis

import (
	"github.com/Ignite-Laboratories/JanOS/core"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

var HomeWorld *World

type World struct {
	Name       string
	Entities   []core.Entity
	Components struct {
		Waveforms WaveformComponents
	}
	Systems []any
	Width   int
	Height  int
}

func (w *World) AddEntity(e core.Entity) {
	w.Entities = append(w.Entities, e)
}

func (w *World) FireBigBang() {
	w.Init()

	ebiten.SetWindowSize(w.Width, w.Height)
	ebiten.SetWindowTitle(w.Name)
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}
}

func (w *World) Init() {
	for _, e := range w.Entities {
		for _, s := range w.Systems {
			if d, ok := s.(core.SystemIniter); ok {
				d.Init(e)
			}
		}
	}
}

func (w *World) Update() error {
	for _, e := range w.Entities {
		for _, s := range w.Systems {
			if u, ok := s.(core.SystemUpdater); ok {
				u.Update(e)
			}
		}
	}
	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for _, e := range w.Entities {
		for _, s := range w.Systems {
			if d, ok := s.(core.SystemDrawer); ok {
				d.Draw(e, screen)
			}
		}
	}

	//ebitenutil.DebugPrint(screen, fmt.Sprintf("Ticks/sec: %0.2f", ebiten.ActualTPS()))
}

func (w *World) Layout(outsideWidth, outsideHeight int) (int, int) {
	return w.Width, w.Height
}
