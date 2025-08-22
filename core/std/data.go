package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
	"github.com/ignite-laboratories/core/sys/support"
	"golang.org/x/exp/slices"
	"sync"
)

// Data is a way to pass slices around in anonymous closures without any worries.  For instance, if you pass a slice as a
// parameter to a function that generates an anonymous method which closed on the slice, appends to the originating slice
// will not persist through to the anonymous method's closure variable.  This is because the anonymous method has closed
// on a COPY of a header pointing to the data, which Go's 'append' method then reassigned outside the anonymous method.
//
// By unifying standardized data in a wrapper structure, you can colloquially guarantee asynchronous access to dimensional
// data by gating append operations through method calls and a mutex - which is exactly what this type provides for you =)
type Data[T any] struct {
	Entity
	mutex  sync.Mutex
	yield  func(uint) []T
	data   []T
	cap    uint
	capped bool
}

// NewData creates a new instance of Data[T].  you may optionally provide a yield function, otherwise calls to Data.Yield
// will return pseudo-random entries from the data set.
func NewData[T any](yieldFn ...func(uint) []T) *Data[T] {
	d := &Data[T]{}

	yield := func(n uint) []T {
		if len(d.data) == 0 {
			return []T{}
		}

		out := make([]T, n)
		for i := uint(0); i < n; i++ {
			out[i] = d.data[num.RandomWithinRange(0, len(d.data)-1)]
		}
		return out
	}
	if len(yieldFn) > 0 {
		yield = yieldFn[0]
	}

	d.yield = yield
	return d
}

// Yield calls an anonymous method that returns a single instance of T 'n' times, appending the results to the underlying slice.
func (d *Data[T]) Yield(n uint) []T {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	out := make([]T, n)

	if d.capped {
		available := int(d.cap) - len(d.data)
		if available < 0 {
			return out
		}

		n = num.Smallest[uint](n, available)
	}

	out = d.yield(n)

	d.mutex.Unlock()
	d.Append(out...)
	return out
}

// Cap sets whether the number of elements should be capped to a provided maximum.
// If 'cap' is true, but no limit is provided, it implies a cap of the current length.
func (d *Data[T]) Cap(cap bool, limit ...uint) {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if cap {
		l := uint(len(d.data))
		if len(limit) > 0 {
			l = limit[0]
		}

		d.cap = l
		d.capped = true
	} else {
		d.capped = false
	}
}

// GetCap returns whether the number of elements has been "capped" and, if so, at what capacity.
//
// NOTE: If the set is uncapped, this will return (false, -1)
func (d *Data[T]) GetCap() (bool, int) {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if !d.capped {
		return false, -1
	}
	return true, int(d.cap)
}

// SetYield overwrites the current yield function with the one provided.
func (d *Data[T]) SetYield(yield func(uint) []T) {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if yield == nil {
		panic("cannot set Data.yield to nil")
	}

	d.yield = yield
}

// Append adds the provided elements in their current order to the end of the underlying slice.
//
// NOTE: If the set is capped, this will apply as many elements as it can from values[0] upward.
//
// For example:
//
//	| A B C D - - - - | ← Original data, capped to 8 positions
//	| Z H G I W K L F | ← Data to append
//	| A B C D Z H G I | ← Resulting set
func (d *Data[T]) Append(values ...T) {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if len(values) == 0 {
		return
	}

	if d.capped {
		available := int(d.cap) - len(d.data)
		length := num.Smallest[uint](len(values), available)

		if length <= 0 {
			return
		}

		values = values[:length]
	}

	d.data = append(d.data, values...)
}

// Prepend adds the provided elements in their current order to the beginning of the underlying slice.
//
// NOTE: If the set is capped, this will apply as many elements as it can from values[len(values)-1] downward.
//
// For example:
//
//	| A B C D - - - - | ← Original data, capped to 8 positions
//	| Z H G I W K L F | ← Data to prepend
//	| W K L F A B C D | ← Resulting set
func (d *Data[T]) Prepend(values ...T) {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if len(values) == 0 {
		return
	}

	if d.capped {
		available := int(d.cap) - len(d.data)
		length := num.Smallest[int](len(values), available)
		subPos := len(d.data) - 1 - length

		if subPos <= 0 {
			return
		}

		values = values[len(values)-available:]
	}

	d.data = append(values, d.data...)
}

// Insert inserts the values at the provided index, pushing the larger indices higher by len(values) positions.
//
// NOTE: If the set is capped, this will push existing elements beyond the cap boundary and trim to length.
//
// For example:
//
//	| A B C D E F G H | ← Original data, capped to 8 positions
//	| - - - K L F - - | ← Data to prepend
//	| A B C K L F D E | ← Resulting set
func (d *Data[T]) Insert(index uint, values ...T) error {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if len(values) == 0 {
		return nil
	}

	if d.capped {
		available := int(d.cap) - len(d.data)
		length := num.Smallest[int](len(values), available)

		if length <= 0 {
			return nil
		}

		values = values[:length]
	}

	length := uint(len(d.data))

	if index > length {
		return d.errorOutOfBounds(index, len(d.data)-1)
	} else if index == length {
		d.data = append(d.data, values...)
		return nil
	}

	d.data = append(d.data[:index], append(values, d.data[index:]...)...)

	if d.capped {
		d.data = d.data[:d.cap]
	}
	return nil
}

// Remove removes the provided indices from the underlying slice.
func (d *Data[T]) Remove(indices ...uint) error {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if len(indices) == 0 {
		return nil
	}

	for _, index := range indices {
		if index >= uint(len(d.data)) {
			return d.errorOutOfBounds(index, len(d.data)-1)
		}
	}

	indices = support.Deduplicate(indices)
	slices.Sort(indices)
	slices.Reverse(indices)

	for _, index := range indices {
		if index+1 == uint(len(d.data)) {
			d.data = d.data[:index]
		} else {
			d.data = append(d.data[:index], d.data[index+1:]...)
		}
	}
	return nil
}

// Clear removes all the elements from the set.
func (d *Data[T]) Clear() {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.data = []T{}
}

// Select returns the requested index range from the underlying slice in the half-open interval [low,high).
// This is meant to mimic the standard slice index accessor pattern in Go, which is also [low,high), allowing
// this type to act as a surrogate 'slice' type.
//
// NOTE: If no 'high' value is provided, a single index is returned.
func (d *Data[T]) Select(low uint, high ...uint) ([]T, error) {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if low >= uint(len(d.data)) {
		return nil, d.errorOutOfBounds(low, len(d.data)-1)
	}

	if len(high) > 0 {
		// Selecting a range
		h := high[0]
		if h > uint(len(d.data)) {
			return nil, d.errorOutOfBounds(h, len(d.data)-1)
		} else if h < low {
			return nil, fmt.Errorf("high index %d must be greater than low index %d", h, low)
		}

		toCopy := d.data[low:h]
		out := make([]T, len(toCopy))
		copy(out, toCopy)

		return out, nil
	}

	// Selecting a single index
	return []T{d.data[low]}, nil
}

// SelectAll returns a copy of the underlying slice.
func (d *Data[T]) SelectAll() []T {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	out := make([]T, len(d.data))
	copy(out, d.data)
	return out
}

// Len returns the length of the underlying slice.
func (d *Data[T]) Len() uint {
	d.SanityCheck()

	d.mutex.Lock()
	defer d.mutex.Unlock()

	return uint(len(d.data))
}

func (d *Data[T]) errorOutOfBounds(index uint, length int) error {
	if length < 0 {
		length = 0
	}

	return fmt.Errorf("index %d out of bounds [0,%d]", index, length)
}

// SanityCheck will validate if the Data has a yield function assigned.  If not, it indicates an erroneous state - causing
// this function to panic.
func (d *Data[T]) SanityCheck() {
	if d.yield == nil {
		panic("Data.yield is nil - please create this type through NewData")
	}
}
