package std

// Movement represents the primitive four degrees of functional traversal. All set traversal operations can be
// distilled down to two actions subdivided into two categories of output -
//
//	     Actions
//	Relatively(𝑛) - Emit along the axis ±𝑛 positions
//	Absolutely(𝑛) - Emit to the fixed position 𝑛 on the axis
//
//	     Output
//	Return a single element
//	Return many elements
//
// While you can write your functions to return however you would like, the convention is:
//
//	Relatively(±𝑛) [either] - Positive movement yields before stepping, negative movement steps and then yields.
//	Absolutely(𝑛) [single] - The target index is always returned
//	Absolutely(𝑛) [many] - Positive movement yields before stepping, negative movement steps and then yields.
//
// See Axis, Emit, Movement, and Pattern
type Movement[T any, TM MovementFn[T]] struct {
	// Relatively moves ±𝑛 positions from the current cursor position.  Positive Movement yields before stepping, negative
	// Movement steps and then yields.
	Relatively TM

	// Absolutely moves directly to position 𝑛 on the shortest path from the current cursor position.  Unless requesting
	// a single position, positive Movement yields before steppingwhile negative Movement steps and then yields.
	Absolutely TM
}
