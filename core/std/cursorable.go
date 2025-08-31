package std

import (
	"core/enum/direction/ordinal"
	"core/sys/num"
	"core/sys/num/bounded"
)

// Cursorable defines the mechanics of a vector traversing a well-defined, yet mutable, data set.
// All bounded.Numeric types are inherently cursors - which naturally includes the components of any Vector type.
// For supporting cursor procedures, please see the contents of the 'core/sys/cursor' package.
//
// There are "six degrees of semantic freedom," broken into two categories -
//
// 'Jump' functions, which jump directly to a target position and yield its element
//
//	Jump( 𝑛 ) 𝑒𝑙𝑒𝑚𝑒𝑛𝑡
//	Jump relatively forward or backwards 𝑛 positions
//
//	JumpTo( 𝑖 ) 𝑒𝑙𝑒𝑚𝑒𝑛𝑡
//	Jump absolutely to position 𝑖
//
//	JumpAlongPath( 𝑝𝑎𝑡ℎ() ) 𝑒𝑙𝑒𝑚𝑒𝑛𝑡
//	Jump absolutely to the result of 𝑝𝑎𝑡ℎ(), which dictates the target, stride, and direction.
//
// 'Walk' functions, which step at a 𝑠𝑡𝑟𝑖𝑑𝑒 interval to a target and yield the found elements.
//
//	Walk( 𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒 ) []𝑒𝑙𝑒𝑚𝑒𝑛𝑡
//	Walk relatively forward or backwards 𝑛 positions
//
//	WalkTo( 𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒, ...𝑑𝑖𝑟𝑒𝑐𝑡𝑖𝑜𝑛 ) []𝑒𝑙𝑒𝑚𝑒𝑛𝑡
//	Walk absolutely to position 𝑖 along the shortest path, or in an optional direction
//
//	WalkAlongPath( 𝑝𝑎𝑡ℎ() ) []𝑒𝑙𝑒𝑚𝑒𝑛𝑡
//	Walk absolutely to the result of 𝑝𝑎𝑡ℎ(), which dictates the target, stride, and direction.
//
// NOTES:
//
// 0. The "shortest path" is defined as the shortest distance towards the target.  In an unclamped cursor, walk operations can take advantage
// of overflowing and underflowing to reach the target quicker - in a clamped cursor, the shortest path is simply the distance between the
// source and target positions.
//
// 1. In walking operations, if the stride doesn't land evenly on the target position, the final "overshooting" step stops at the target position.
//
// 2. While I use the term '𝑠𝑡𝑟𝑖𝑑𝑒' for its conciseness, it's better to consider this as a way of defining the 𝑟𝑒𝑠𝑜𝑙𝑢𝑡𝑖𝑜𝑛 of an infinitely bounded space. The
// closed interval of [0.0, 1.0] is infinitely resolute - but the stride defines which resolution you wish to perceive it as.  Just remember, the
// IEEE 754 floating point specification used by Go's float32 and float64 types has limitations for floating point operations!  If you keep your
// stride to a reasonable 𝑟𝑒𝑠𝑜𝑙𝑢𝑡𝑖𝑜𝑛, you'll be a-okay =)
//
// 3. All letter vectors inherently are cursors, where their numeric type IS the cursor's data type!  The Cursor type is usually used to navigate slices of
// 1D data, but any letter vector can be used against an external multidimensional data set.  When cursoring a letter vector through its addressable range,
// the yielded elements are the positions the vector visited - allowing you to translate those steps into the elements from your dataset.
//
// See Cursorable and Cursor
type Cursorable[T num.Primitive, TBounded bounded.Numeric[T]] interface {
	Jump(n T) T
	JumpTo(i T) T
	JumpAlongPath(bounded.PathFn[T, TBounded]) T

	Walk(n T, stride T) []T
	WalkTo(i T, stride T, direction ...ordinal.Direction) []T
	WalkAlongPath(path bounded.PathFn[T, TBounded], stride T, direction ...ordinal.Direction) []T
}
