package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/host/mouse"
	"time"
)

var framerate = 60.0 //hz
var xTimeScale = std.TimeScale[int]{Duration: time.Second * 2, Height: 3640}

var xCoords = temporal.Calculation(core.Impulse, when.Frequency(&mouse.SampleRate), false, SampleX)

func main() {
	viewport.NewBasicWaveform(true, when.Frequency(&framerate), "Mouse X", nil, nil, &xTimeScale, false, xCoords)
	core.Impulse.Spark()
}

func SampleX(ctx core.Context) int {
	return mouse.Sample().Position.X
}
