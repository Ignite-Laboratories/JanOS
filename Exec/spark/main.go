package main

import (
	"github.com/Ignite-Laboratories/JanOS/Spark"
	"github.com/Ignite-Laboratories/JanOS/Spark/Systems"
	"github.com/hajimehoshi/ebiten/v2"
)

var AssetsToLoad = map[string]string{
	"sine.1k":     "audio\\sine.1k.wav",
	"segoe-print": "fonts\\segoepr.ttf",
}

func main() {
	Spark.Universe = &Spark.World{
		Assets: Spark.NewAssetManager("c:\\source\\ignite\\janos\\assets", AssetsToLoad),
		Systems: []Spark.System{
			Systems.NewAudioSystem(),
			Systems.NewInputSystem(),
			Systems.NewWaveformVisualizerSystem(),
		},
	}

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
