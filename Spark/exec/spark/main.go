package main

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/Ignite-Laboratories/JanOS/Logic/Assets"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

/**
SETUP
*/

const (
	screenWidth  = 1024
	screenHeight = 768
)

var assetSystem = &Assets.AssetSystem{
	BaseDirectory: "c:\\source\\ignite\\janos\\Assets",
	ToLoad: map[string]string{
		"sine.1k":     "audio\\sine.1k.wav",
		"segoe-print": "fonts\\segoepr.ttf",
	},
}

/**
RUN
*/

var world *Logic.World

func main() {
	game := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Spark")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

/**
SUPPORT
*/

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
