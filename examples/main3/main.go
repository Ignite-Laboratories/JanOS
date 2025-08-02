package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/host/mouse"
	"math"
)

func init() {
	//mouse.Coordinates.Unmute()
	//temporal.Analyzer[std.XY[int], any, any](core.Impulse, when.EighthSpeed(&mouse.SampleRate), false, Print, mouse.Coordinates)
	mouse.Reaction(core.Impulse, std.Ref(2048.0), Velocity)
	mouse.Reaction(core.Impulse, &mouse.SampleRate, Feedback)
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

func Feedback(ctx core.Context, old std.Data[std.MouseState], current std.Data[std.MouseState]) {
	if current.Point.Position.X > 1024 {
		mouse.SampleRate = 2048.0
	} else {
		mouse.SampleRate = 2.0
	}
	fmt.Println(current.Point.Position)
}

func Print(ctx core.Context, cache *any, data []std.Data[std.MouseState]) any {
	points := make([]std.XY[int], len(data))
	for i, v := range data {
		points[i] = v.Point.Position
	}
	fmt.Printf("[%d] %v\n", ctx.Beat, points)
	return nil
}
