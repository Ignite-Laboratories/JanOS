// Package direction provides access to the Axis interface.
package direction

import (
	"core/enum/direction/awareness"
	"core/enum/direction/cardinal"
	"core/enum/direction/consciousness"
	"core/enum/direction/ordinal"
	"core/enum/direction/orthogonal"
	"core/enum/direction/temporal"
	"core/enum/direction/transmittal"
	"core/enum/direction/universal"
)

// Any represents general directionality and includes both cardinal and abstract reference points in space and time.
//
// All dimensions can be distilled down to an infinitely repeating number line which can be traversed in binary directions -
// but, as you layer these dimensions on top of each other, they orthographically align relative to one another.  The terminology
// used to describe this is entirely dependent upon context, and as such I've provided a robust set of general abstract dimensions
// from which to describe this mechanic in code.  It truly does NOT matter which you use, as long as the called method knows
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
//	orthogonal.Up: Towards the top of the orthogonal XY plane
//	orthogonal.Down: Towards the bottom of the orthogonal XY plane
//	orthogonal.Left: Towards the left of the orthogonal XY plane
//	orthogonal.Right: Towards the right of the orthogonal XY plane
//	orthogonal.In: Negatively along the Z axis perpendicular to the orthogonal XY plane
//	orthogonal.Out: Positively along the Z axis perpendicular to the orthogonal XY plane
//
//	ordinal.Negative:  i - 1
//	ordinal.Current: i + 0
//	ordinal.Positive:   i + 1
//
//	temporal.Future: Anticipation
//	temporal.Present: Experience
//	temporal.Past: Reflection
//
//	transmittal.Inbound: Receiving
//	transmittal.Outbound: Transmitting
//	transmittal.Bidirectional: Discourse
//
// In addition, several "higher-order" dimensional descriptions are defined as well.  These are entirely abstract and imply
// no logical order of their sub-dimensions, only ordinal terminology for traversing them.
//
//	awareness.Nascent: Obliviousness
//	awareness.Naive: Willingness
//	awareness.Mature: Understanding
//
//	consciousness.Ignorant: Selfishness
//	consciousness.Emergent: Self-Awareness
//	consciousness.Aware: Selflessness
//
//	reality.Chaos: Disorder
//	reality.Coherence: Identification
//	reality.Stability: Alignment
//
// See direction.Any, cardinal.Direction, orthogonal.Direction, ordinal.Direction, temporal.Direction, transmittal.Direction, awareness.Direction, consciousness.Direction, universal.Direction,
// Spatial, SpaceTime, Awareness, Consciousness, and Universal
type Any interface {
	cardinal.Direction | orthogonal.Direction | ordinal.Direction | temporal.Direction | transmittal.Direction | awareness.Direction | consciousness.Direction | universal.Direction
}

// Ordinal represents the logical directional order of elements.
//
// See direction.Any and ordinal.Direction
type Ordinal interface {
	ordinal.Direction
}

// Spatial represents any traversable axis of XYZ space.
//
// See direction.Any and orthogonal.Direction
type Spatial interface {
	orthogonal.Volume
}

// SpaceTime represents any traversable axis of space or time.
//
// See direction.Any, orthogonal.Volume, and temporal.Direction
type SpaceTime interface {
	orthogonal.Volume | temporal.Direction
}

// Awareness represents any traversable axis through the 5th dimension.
//
// See direction.Any, SpaceTime, and awareness.Direction
type Awareness interface {
	SpaceTime | awareness.Direction
}

// Consciousness represents any traversable axis through the 6th dimension.
//
// See direction.Any, SpaceTime, awareness.Direction, and consciousness.Direction
type Consciousness interface {
	SpaceTime | awareness.Direction | consciousness.Direction
}

// Universal represents any traversable axis through the 7th dimension.
//
// See direction.Any, SpaceTime, awareness.Direction, consciousness.Direction, and universal.Direction
type Universal interface {
	SpaceTime | awareness.Direction | consciousness.Direction | universal.Direction
}
