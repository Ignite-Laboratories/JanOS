package std

import (
	"core/enum/direction/ordinal"
	"core/sys/num"
)

// A PathFn should plot a course for a Cursor to Step along.
//
// See Cursor, PathFn, Step, and num.Breach.
type PathFn[T num.Advanced] func(T) []Step[T]

// A Step indicates the Target index and Direction a cursor should step to.  In addition, the amount this operation
// should num.Breach the boundaries is precalculated.  This can be ignored, or it could be used for validation that your
// movement operation matched what the plotted course desired.
//
// See Cursor, PathFn, Step, and num.Breach.
type Step[T num.Advanced] struct {
	Direction ordinal.Direction
	Target    T
	Breach    num.Breach
}

// A Cursor entity is one that can be moved through an abstract space via "jump" or "walk" operations.
// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
// the path.
//
// NOTE: Each movement operation may optionally output its num.Breach - the extent to which it exceeded the bounds.
// This parameter may be ignored if they're irrelevant to you.
//
// These operations are broken into "six degrees of semantic freedom"
//
// - Jump(𝑛) - Relatively move by 𝑛 and yield
//
// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
//
// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
//
// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
//
// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
//
// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way to each position.
//
// See Cursor, PathFn, Step, and num.Breach.
type Cursor[T num.Advanced] interface {
	// A Cursor entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - the extent to which it exceeded the bounds.
	// This parameter may be ignored if they're irrelevant to you.
	//
	// "The six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way to each position.
	//
	// See Cursor, PathFn, Step, and num.Breach.
	Jump(n T, out ...*num.Breach) T

	// A Cursor entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - the extent to which it exceeded the bounds.
	// This parameter may be ignored if they're irrelevant to you.
	//
	// "The six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way to each position.
	//
	// See Cursor, PathFn, Step, and num.Breach.
	JumpTo(i T, out ...*num.Breach) T

	// A Cursor entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - the extent to which it exceeded the bounds.
	// This parameter may be ignored if they're irrelevant to you.
	//
	// "The six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way to each position.
	//
	// See Cursor, PathFn, Step, and num.Breach.
	JumpAlong(path PathFn[T], out ...*num.Breach) []T

	// A Cursor entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - the extent to which it exceeded the bounds.
	// This parameter may be ignored if they're irrelevant to you.
	//
	// "The six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way to each position.
	//
	// See Cursor, PathFn, Step, and num.Breach.
	Walk(n T, stride T, out ...*num.Breach) []T

	// A Cursor entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - the extent to which it exceeded the bounds.
	// This parameter may be ignored if they're irrelevant to you.
	//
	// "The six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way to each position.
	//
	// See Cursor, PathFn, Step, and num.Breach.
	WalkTo(i T, stride T, direction ordinal.Direction, out ...*num.Breach) []T

	// A Cursor entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - the extent to which it exceeded the bounds.
	// This parameter may be ignored if they're irrelevant to you.
	//
	// "The six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒 along the way to each position.
	//
	// See Cursor, PathFn, Step, and num.Breach.
	WalkAlong(path PathFn[T], stride T, out ...*[]num.Breach) []T
}
