package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/Ignite-Laboratories/JanOS/systems"
	"image/color"
)

type myScene struct{}

func (*myScene) Type() string { return "JanOS" }

func (*myScene) Preload() {
}

func (*myScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	common.SetBackground(color.White)

	engo.Input.RegisterButton("Analyze", engo.KeyF1)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.AudioSystem{})
	world.AddSystem(&common.MouseSystem{})

	world.AddSystem(&systems.AssetSystem{
		FilesToLoad: map[string]string{
			"sine.1k": "audio\\sine.1k.wav",
		},
	})
	world.AddSystem(&systems.WaveformSystem{})
}

func main() {
	opts := engo.RunOptions{
		Title:          "JanOS",
		Width:          640,
		Height:         480,
		StandardInputs: true,
	}
	engo.Run(opts, &myScene{})
}
