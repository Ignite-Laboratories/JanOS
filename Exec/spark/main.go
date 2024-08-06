package main

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/Ignite-Laboratories/JanOS/Logic/Systems"
	"github.com/hajimehoshi/ebiten/v2"
)

var AssetsToLoad = map[string]string{
	"sine.1k":     "audio\\sine.1k.wav",
	"segoe-print": "fonts\\segoepr.ttf",
}

var world = &Logic.World{
	Assets: Logic.NetAssetManager("c:\\source\\ignite\\janos\\assets", AssetsToLoad),
	Systems: []Logic.System{
		Systems.NewAudioSystem(),
		Systems.NewInputSystem(),
		Systems.NewWaveformVisualizerSystem(),
	},
}

func main() {
	game := Logic.Game{
		WindowTitle:  "Spark",
		ScreenWidth:  1024,
		ScreenHeight: 768,
		World:        world,
		OnUpdate: func() {

		},
		OnDraw: func(screen *ebiten.Image) {
			for _, s := range world.Systems {
				if d, ok := s.(Logic.SystemDrawer); ok {
					d.Draw(screen)
				}
			}
		},
	}
	game.Run()
}
