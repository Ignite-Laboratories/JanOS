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

func main() {
	game := Logic.Game{
		WindowTitle:  "Spark",
		ScreenWidth:  1024,
		ScreenHeight: 768,
		World: &Logic.World{
			Systems: []Logic.System{
				Systems.NewAssetSystem("c:\\source\\ignite\\janos\\Assets", AssetsToLoad),
				Systems.NewAudioSystem(),
				Systems.NewInputSystem(),
			},
		},
		OnUpdate: func() {

		},
		OnDraw: func(screen *ebiten.Image) {

		},
	}
	game.Run()
}
