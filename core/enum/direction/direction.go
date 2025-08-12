// Package direction provides access to the Direction interface.
package direction

import (
	"github.com/ignite-laboratories/core/enum/direction/cardinal"
	"github.com/ignite-laboratories/core/enum/direction/ordinal"
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/enum/direction/temporal"
	"github.com/ignite-laboratories/core/enum/direction/traffic"
	"github.com/ignite-laboratories/core/std/num"
)

// Any represents general directionality and includes both cardinal and abstract reference points in time and space.
//
// All dimensions can be distilled down to a number line which can be traversed in binary directions - but, as you layer
// dimensions on top of each other, they orthographically align relative to one another.  The terminology used to describe
// this is entirely dependent upon context, and as such I've provided a robust set of general abstract dimensions from
// which to describe this mechanic in code.  It truly does NOT matter which you use, as long as the called method knows
// how to talk in THAT language. =)
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, you walk "latitudinally" between rows along the Y axis and "longitudinally" between columns along the X axis.  Against a voxel,
// you'd walk negatively "in" or positively "out" along the Z axis.
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation =)
//
// The sub-directions each have an explicitly defined purpose relative to a fixed point -
//
// NOTE: Like-typed directions can be combined to form complex directions.
//
//	cardinal.South: Calculation
//	cardinal.North: Accumulation
//	cardinal.West: Scale
//	cardinal.East: Reduction
//
//	ordinal.Before:  i - 1
//	ordinal.Current: i + 0
//	ordinal.After:   i + 1
//
//	temporal.Future: Anticipation
//	temporal.Present: Experience
//	temporal.Past: Reflection
//
//	traffic.Inbound: Receiving
//	traffic.Outbound: Transmitting
//	traffic.Bidirectional: Discourse
//
//	orthogonal.Up: Towards the top of the orthogonal XY plane
//	orthogonal.Down: Towards the bottom of the orthogonal XY plane
//	orthogonal.Left: Towards the left of the orthogonal XY plane
//	orthogonal.Right: Towards the right of the orthogonal XY plane
//	orthogonal.In: Negatively along the Z axis perpendicular to the orthogonal XY plane
//	orthogonal.Out: Positively along the Z axis perpendicular to the orthogonal XY plane
//
// See direction.Any, cardinal.Direction, orthogonal.Direction, ordinal.Direction, temporal.Direction, traffic.Direction
type Any[T num.Primitive] interface {
	cardinal.Direction[T] | orthogonal.Direction[T] | ordinal.Direction[T] | temporal.Direction[T] | traffic.Direction[T]
}
