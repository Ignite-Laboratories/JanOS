package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/host/mouse"
	"math"
)

func init() {
	mouse.Reaction(core.Impulse, &mouse.SampleRate, Velocity)
}

func main() {
	core.Impulse.Spark()
}

func Velocity(ctx core.Context, old std.Data[std.MouseState], current std.Data[std.MouseState]) {
	delta := current.Point.Position.X - old.Point.Position.X
	deltaAbs := math.Abs(float64(delta))
	if deltaAbs > 100 {
		fmt.Println("Slow down!")
	}
}
