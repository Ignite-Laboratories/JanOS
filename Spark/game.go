package Spark

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

var Universe *World

type Game struct {
	WindowTitle   string
	ScreenWidth   int
	ScreenHeight  int
	OnUpdate      func()
	PreDraw       func(screen *ebiten.Image)
	PostDraw      func(screen *ebiten.Image)
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
	// On the first tick, we initialize all the OLD
	if !g.isInitialized {
		Universe.Messaging = NewNexus()
		Universe.Assets.Initialize()
		for _, system := range Universe.Systems {
			log.Printf("%s System Initializing", system.GetName())
			system.Initialize()
			log.Printf("%s System Initialized", system.GetName())
		}
		g.isInitialized = true
	} else {
		// On subsequent ticks, we fire the tick function
		for _, system := range Universe.Systems {
			messages := Universe.GetMessages(system.GetEntity())
			system.Tick(messages)
		}
	}

	g.OnUpdate()

	// Update the message queue last
	Universe.Messaging.Cycle()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.PreDraw(screen)
	for _, e := range Universe.Entities {
		for _, s := range Universe.Systems {
			if d, ok := s.(SystemDrawer); ok {
				d.OnDraw(e, screen)
			}
		}
	}
	g.PostDraw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
