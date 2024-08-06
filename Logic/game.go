package Logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	WindowTitle   string
	ScreenWidth   int
	ScreenHeight  int
	Nexus         *Nexus
	World         *World
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
	// On the first tick, we initialize all the OLD
	if !g.isInitialized {
		g.World.Messaging = NewNexus()
		g.World.Assets.Initialize(g.World)
		for _, system := range g.World.Systems {
			log.Printf("%s System Initializing", system.GetName())
			system.Initialize(g.World)
			log.Printf("%s System Initialized", system.GetName())
		}
		g.isInitialized = true
	} else {
		// On subsequent ticks, we fire the tick function
		for _, system := range g.World.Systems {
			messages := g.World.GetMessages(system.GetEntity())
			system.Tick(g.World, messages)
		}
	}

	g.OnUpdate()

	// Update the message queue last
	g.World.Messaging.Cycle()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.OnDraw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
