package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	mouse2 "github.com/ignite-laboratories/core/std/mouse"
	"github.com/ignite-laboratories/core/sys/atlas"
	"github.com/ignite-laboratories/host/mouse"
	"math"
)

func init() {
	mouse.Reaction(atlas.Impulse, &mouse.SampleRate, Velocity)
}

func main() {
	atlas.Impulse.Spark()
}

func Velocity(ctx core.Context, old std.Data[mouse2.State], current std.Data[mouse2.State]) {
	delta := current.Point.Position.X - old.Point.Position.X
	deltaAbs := math.Abs(float64(delta))
	if deltaAbs > 100 {
		fmt.Println("Slow down!")
	}
}
