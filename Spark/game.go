package Spark

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	WindowTitle   string
	ScreenWidth   int
	ScreenHeight  int
	World         *Logic.World
	OnUpdate      func()
	OnDraw        func(screen *ebiten.Image)
	isInitialized bool
}

func (g *Game) Run() {
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle(g.WindowTitle)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	// On the first tick, we initialize all the systems
	if !g.isInitialized {
		for _, system := range g.World.Systems {
			system.Initialize(g.World)
		}
		g.isInitialized = true
	} else {
		// On subsequent ticks, we fire the tick function
		for _, system := range g.World.Systems {
			system.Tick(g.World)
		}
	}

	g.OnUpdate()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.OnDraw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
