package main

import (
	"github.com/Ignite-Laboratories/JanOS/Logic"
	"github.com/Ignite-Laboratories/JanOS/Logic/Systems"
	"github.com/Ignite-Laboratories/JanOS/Spark"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := Spark.Game{
		WindowTitle:  "Spark",
		ScreenWidth:  1024,
		ScreenHeight: 768,
		World: &Logic.World{
			Systems: []Logic.System{
				Systems.NewAssetSystem("c:\\source\\ignite\\janos\\Assets", map[string]string{
					"sine.1k":     "audio\\sine.1k.wav",
					"segoe-print": "fonts\\segoepr.ttf",
				}),
			},
		},
		OnUpdate: func() {

		},
		OnDraw: func(screen *ebiten.Image) {

		},
	}
	game.Run()
}
