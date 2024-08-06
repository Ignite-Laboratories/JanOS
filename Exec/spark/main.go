package main

import (
	"github.com/Ignite-Laboratories/JanOS/Spark"
	"github.com/Ignite-Laboratories/JanOS/Spark/Systems"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

var AssetsToLoad = map[string]string{
	"sine.1k":     "audio\\sine.1k.wav",
	"segoe-print": "fonts\\segoepr.ttf",
}

func main() {
	cursoring := Systems.NewCursoringSystem()
	oscillation := Systems.NewOscillationSystem()

	Spark.Universe = &Spark.World{
		Assets: Spark.NewAssetManager("c:\\source\\ignite\\janos\\assets", AssetsToLoad),
		Systems: []Spark.System{
			oscillation,
			cursoring,
			Systems.NewWaveformVisualizerSystem(),
		},
	}

	o1 := oscillation.StartOscillator(1, 1, time.Duration(10)*time.Second)
	oscillation.StartOscillator(1, 1, time.Duration(5)*time.Second)
	oscillation.StartOscillator(1, 1, time.Duration(1)*time.Second)
	oscillation.StartOscillator(1, 1, time.Duration(5)*time.Second)
	oscillation.StartOscillator(1, 1, time.Duration(10)*time.Second)

	cursoring.StartCursor(o1, 100, time.Duration(1)*time.Second)

	game := Spark.Game{
		WindowTitle:  "Spark",
		ScreenWidth:  1024,
		ScreenHeight: 768,
		OnUpdate: func() {
		},
		PreDraw: func(screen *ebiten.Image) {
		},
		PostDraw: func(screen *ebiten.Image) {
		},
	}
	game.Run()
}
