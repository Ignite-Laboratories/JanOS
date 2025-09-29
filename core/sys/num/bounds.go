package num

// Bounds allows you to configure a 𝑡𝑖𝑛𝑦 mathematical operation.  All fields are nilable and only those with a non-nil
// value will be included in the calculation.
//
// NOTE: The values are dereferenced before the chain of calculation begins to prevent race conditions.
type Bounds struct {
	Minimum any
	Maximum any
	Clamp   *bool
}
