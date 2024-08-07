package main

import (
	"github.com/Ignite-Laboratories/JanOS/Spark"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

var AssetsToLoad = map[string]string{
	"sine.1k":     "audio\\sine.1k.wav",
	"segoe-print": "fonts\\segoepr.ttf",
}

func main() {
	cursoring := Spark.NewCursoringSystem()
	oscillation := Spark.NewOscillationSystem()
	waveform := Spark.NewWaveformVisualizerSystem()

	Spark.Universe = &Spark.World{
		Assets:                Spark.NewAssetManager("c:\\source\\ignite\\janos\\Assets", AssetsToLoad),
		Systems:               []Spark.System{cursoring, oscillation, waveform},
		Cursoring:             cursoring,
		Oscillation:           oscillation,
		WaveformVisualization: waveform,
	}

	o1 := oscillation.StartOscillator(1, 1, time.Duration(1)*time.Second)
	c1 := cursoring.StartCursor(o1, 100, time.Duration(1)*time.Second)
	waveform.Visualize(c1)

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
