package semantic

// Set represents any semantic.Slice of elements and the semantics of how to express data from them.
//
//
//
// NOTE: 'context' is a simplified alias for Context.
type Set[T any] interface {
	*Slice[T] | *Unique[T] | *context[T] | *Experience[T]

	// Cap should set whether the number of elements should be capped to a provided maximum.
	// If no value is provided, it should imply a cap of the current length.
	Cap(bool, ...uint)

	// GetCap should indicate if the set is capped and, if so, what the cap value is.  If the data is uncapped, the
	// returned value should be (false, -1)
	GetCap() (bool, int)

	// Yield should call 'fn' 'n' times, then record and return the yielded elements.
	Yield(n uint, fn func() []T) []T

	// Len should return the number of elements within the set.
	Len() uint

	// Append should add the provided elements to the end of the set in the order received from left→to→right.
	Append(...T)

	// Prepend should add the provided elements to the beginning of the set in the order received from left→to→right.
	Prepend(...T)

	// Clear should remove all elements from the set.
	Clear()

	// Insert should insert the provided elements into the set, or error if the provided index position is out of bounds.
	Insert(uint, ...T) error

	// Remove should remove the provided indices from the set.  If any of the requested indices are out of bounds,
	// the entire operation should not commit any changes.
	//
	// NOTE: If no indices are provided, this should do nothing.  For clearing the entire set, see the Clear operation.
	Remove(...uint) error
	Select(uint, ...uint) ([]T, error)
	SelectAll() []T
	SanityCheck()
}
