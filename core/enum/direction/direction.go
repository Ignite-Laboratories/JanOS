// Package direction provides access to the Direction interface.
package direction

import (
	"github.com/ignite-laboratories/core/enum/direction/cardinal"
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/enum/direction/relatively"
	"github.com/ignite-laboratories/core/enum/direction/temporal"
	"github.com/ignite-laboratories/core/enum/direction/traffic"
)

// Direction represents general directionality and includes both cardinal and abstract reference points in time and space.
//
// Abstractly, the result of calculation (the target) is always relatively "down" (or "towards the enemy gate") no matter YOUR orientation
// in space.  Mentally this may be the direction of "gravity" while standing up and writing calculations on a whiteboard, but I think Ender
// described it best.  All binary data is oriented with the most-significant side towards the "left" (or "west").  When operating against a
// matrix, the abstract orientation of "rows" are aligned with the cardinal direction of "south".
//
// Abstract references consider your relative orientation as you float through the void of time and spatial calculation.
//
// The sub-directions each have an explicitly defined purpose relative to a fixed point -
//
// cardinal.Direction
//
//	cardinal.South: Calculation
//	cardinal.North: Accumulation
//	cardinal.West: Scale
//	cardinal.East: Reduction
//
// temporal.Direction
//
//	temporal.Future: Anticipation
//	temporal.Present: Experience
//	temporal.Past: Reflection
//
// traffic.Direction
//
//	traffic.Inbound: Receiving
//	traffic.Outbound: Transmitting
//	traffic.Bidirectional: Communication
//
// orthogonal.Direction
//
//	orthogonal.Up: Towards the top of the viewport
//	orthogonal.Down: Towards the bottom of the viewport
//	orthogonal.Left: Towards the left of the viewport
//	orthogonal.Right: Towards the right of the viewport
//	orthogonal.In: Towards the viewport
//	orthogonal.Out: Away from the viewport
//
// relatively.Direction
//
//	relatively.Before:  i - 1 NOTE: This is hardcoded as -1
//	relatively.Current: i + 0 NOTE: This is hardcoded as  0
//	relatively.After:   i + 1 NOTE: This is hardcoded as  1
type Direction interface {
	cardinal.Direction | orthogonal.Direction | relatively.Relative | temporal.Direction | traffic.Direction
}
