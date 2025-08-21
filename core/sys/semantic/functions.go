package semantic

// YieldFn is the central pillar of raster calculation.  While the concept is best described from a one-dimensional slice,
// bear in mind that T can also be a multidimensional slice!  The concept grows as you define more complex yield
// functions, but fundamentally remains the same at every single level.  Record functions take in an incrementing value
// which indicates what step in the calculation they are currently performing and the slice from which to perform their
// calculation.  They are fully permitted to mutate the slice to whatever degree they desire in the process, and whatever
// result they deliver is what will get recorded as the new 'set' for the next yield operation to work against.  If they
// wish to deliver "artifacts" to the caller, they are permitted to do so through the 'artifacts' output.
//
// Why?  Well, this function allows many different operations - defined by the standard yield functions under the
// semantic.Fn structure:
//
// slice.Selection(low, ...high)
// slice.RandomSelection(...count)
// slice.All()
// slice.Reverse()
// slice.First()
// slice.Last()
// slice.Append(data)
// slice.Prepend(data)
// slice.Insert(index, data)
// slice.Remove(indices)
//
// slice.Slice.Append(func() ([]T, bool))
// slice.Slice.Prepend(func() ([]T, bool))
// slice.Slice.Insert(index, func() ([]T, bool))
//
// random.FromSet(count)
// random.Bounded(record, count, low, high)
// random.BoundedByType[T](record, count)
type YieldFn[T any] func(slice Slice[T]) (mutated Slice[T], artifacts Slice[T])

// Cap should set whether the number of elements should be capped to a provided maximum.
// If no value is provided, it should imply a cap of the current length.
Cap(bool, ...uint)

// GetCap should indicate if the set is capped and, if so, what the cap value is.  If the data is uncapped, the
// returned value should be (false, -1)
GetCap() (bool, int)

// Record should call 'fn' 'n' times, then record and return the yielded elements.
Yield(n uint, fn func() []T) []T

// Len should return the number of elements within the set.
Len() uint

// Clear should remove all elements from the set.
Clear()

