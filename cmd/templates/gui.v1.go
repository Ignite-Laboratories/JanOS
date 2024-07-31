package templates

import (
	"fmt"
	"github.com/Ignite-Laboratories/JanOS/entities"
	"github.com/Ignite-Laboratories/JanOS/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

const (
	tileSize     = 64
	screenWidth  = tileSize * 30
	screenHeight = tileSize * 20
)

type World struct {
	Entities   []entities.Entity
	Components struct {
	}
	Systems []any
}

func (g *World) Init() {
	for _, e := range g.Entities {
		for _, s := range g.Systems {
			if d, ok := s.(systems.SystemIniter); ok {
				d.Init(e)
			}
		}
	}
}

func (g *World) Update() error {
	for _, e := range g.Entities {
		for _, s := range g.Systems {
			if u, ok := s.(systems.SystemUpdater); ok {
				u.Update(e)
			}
		}
	}
	return nil
}

func (g *World) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for _, e := range g.Entities {
		for _, s := range g.Systems {
			if d, ok := s.(systems.SystemDrawer); ok {
				d.Draw(e, screen)
			}
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Ticks/sec: %0.2f", ebiten.ActualTPS()))
}

func (g *World) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	world := &World{}

	// Setup the Systems

	inputSystem := &systems.InputSystem{}

	world.Systems = []any{
		inputSystem,
	}

	world.Init()

	ebiten.SetWindowSize(screenWidth*4, screenHeight*4)
	ebiten.SetWindowTitle("JanOS")
	if err := ebiten.RunGame(world); err != nil {
		log.Fatal(err)
	}
}
