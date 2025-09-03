package std

import (
	"core/enum/direction/ordinal"
	"core/sys/num"
)

// A PathFn should plot a course for a Cursorable to Step along.
//
// See Cursorable, CursorState, PathFn, and Step
type PathFn[T num.Advanced] func(Cursorable[T]) []Step[T]

// A Step indicates the Target index and Direction a cursor should step to.  In addition, the amount this operation
// should num.Breach the boundaries is precalculated.  This can be ignored, or it could be used for validation that your
// movement operation matched what the plotted course desired.
//
// See Cursorable, CursorState, PathFn, and Step
type Step[T num.Advanced] struct {
	Direction ordinal.Direction
	Target    T
	Breach    num.Breach
}

// CursorState indicates the current state of the Cursorable.
//
// Bounded - indicates if the num.Advanced value is 'bounded' beyond the confines of its underlying type.
//
// Clamped - indicates if the num.Advanced value is 'clamped' (meaning it does not allow overflow/underflow of its bounds).
//
// Minimum and Maximum - indicate the num.Advanced's current boundaries (this can be ignored for unbounded values).
//
// See Cursorable, CursorState, PathFn, and Step
type CursorState[T num.Advanced] struct {
	Position T

	Bounded bool
	Clamped bool
	Minimum T
	Maximum T
}

// A Cursorable entity is one that can be moved through an abstract space via "jump" or "walk" operations.
// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
// the path.  To get the current position and boundary details of the cursor, please use the State() method.
//
// NOTE: Each movement operation may optionally output its num.Breach - or, "the extent to which it breached the bounds."
// This parameter may be ignored if it's irrelevant to you.
//
// These operations are broken into "six degrees of semantic freedom"
//
// - Jump(𝑛) - Relatively move by 𝑛 and yield
//
// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
//
// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
//
// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
//
// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
//
// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn while yielding at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
//
// - State() - Returns the current CursorState - which includes the cursor's position.
//
// See Cursorable, CursorState, PathFn, and Step
type Cursorable[T num.Advanced] interface {
	// A Cursorable entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.  To get the current position and boundary details of the cursor, please use the State() method.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - or, "the extent to which it breached the bounds."
	// This parameter may be ignored if it's irrelevant to you.
	//
	// These operations are broken into "six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn while yielding at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - State() - Returns the current CursorState - which includes the cursor's position.
	//
	// See Cursorable, CursorState, PathFn, and Step
	State() CursorState[T]

	// A Cursorable entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.  To get the current position and boundary details of the cursor, please use the State() method.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - or, "the extent to which it breached the bounds."
	// This parameter may be ignored if it's irrelevant to you.
	//
	// These operations are broken into "six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn while yielding at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - State() - Returns the current CursorState - which includes the cursor's position.
	//
	// See Cursorable, CursorState, PathFn, and Step
	Jump(n T, out ...*num.Breach) T

	// A Cursorable entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.  To get the current position and boundary details of the cursor, please use the State() method.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - or, "the extent to which it breached the bounds."
	// This parameter may be ignored if it's irrelevant to you.
	//
	// These operations are broken into "six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn while yielding at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - State() - Returns the current CursorState - which includes the cursor's position.
	//
	// See Cursorable, CursorState, PathFn, and Step
	JumpTo(i T, out ...*num.Breach) T

	// A Cursorable entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.  To get the current position and boundary details of the cursor, please use the State() method.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - or, "the extent to which it breached the bounds."
	// This parameter may be ignored if it's irrelevant to you.
	//
	// These operations are broken into "six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn while yielding at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - State() - Returns the current CursorState - which includes the cursor's position.
	//
	// See Cursorable, CursorState, PathFn, and Step
	JumpAlong(path PathFn[T], out ...*num.Breach) []T

	// A Cursorable entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.  To get the current position and boundary details of the cursor, please use the State() method.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - or, "the extent to which it breached the bounds."
	// This parameter may be ignored if it's irrelevant to you.
	//
	// These operations are broken into "six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn while yielding at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - State() - Returns the current CursorState - which includes the cursor's position.
	//
	// See Cursorable, CursorState, PathFn, and Step
	Walk(n T, stride T, out ...*num.Breach) []T

	// A Cursorable entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.  To get the current position and boundary details of the cursor, please use the State() method.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - or, "the extent to which it breached the bounds."
	// This parameter may be ignored if it's irrelevant to you.
	//
	// These operations are broken into "six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn while yielding at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - State() - Returns the current CursorState - which includes the cursor's position.
	//
	// See Cursorable, CursorState, PathFn, and Step
	WalkTo(i T, stride T, direction ordinal.Direction, out ...*num.Breach) []T

	// A Cursorable entity is one that can be moved through an abstract space via "jump" or "walk" operations.
	// A "jump" operation yields the jumped-to element, while a "walk" operation yields many elements along
	// the path.  To get the current position and boundary details of the cursor, please use the State() method.
	//
	// NOTE: Each movement operation may optionally output its num.Breach - or, "the extent to which it breached the bounds."
	// This parameter may be ignored if it's irrelevant to you.
	//
	// These operations are broken into "six degrees of semantic freedom"
	//
	// - Jump(𝑛) - Relatively move by 𝑛 and yield
	//
	// - JumpTo(𝑖) - Absolutely move to position 𝑖 and yield
	//
	// - JumpAlong(𝑝𝑎𝑡ℎ) - Yield at each Step plotted by the provided PathFn.
	//
	// - Walk(𝑛, 𝑠𝑡𝑟𝑖𝑑𝑒) - Relatively move by 𝑛 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkTo(𝑖, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to position 𝑖 and yield at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - WalkAlong(𝑝𝑎𝑡ℎ, 𝑠𝑡𝑟𝑖𝑑𝑒) - Absolutely move to each Step plotted by the provided PathFn while yielding at a given 𝑠𝑡𝑟𝑖𝑑𝑒.
	//
	// - State() - Returns the current CursorState - which includes the cursor's position.
	//
	// See Cursorable, CursorState, PathFn, and Step
	WalkAlong(path PathFn[T], stride T, out ...*[]num.Breach) []T
}
