package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/std/num"
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
	mutex sync.Mutex
	yield func() T
	data  []T
}

// NewData creates a new instance of Data[T].  you may optionally provide a yield function, otherwise calls to Data.Yield
// will return a random entry from the data set.
func NewData[T any](yieldFn ...func() T) *Data[T] {
	d := &Data[T]{}

	yield := func() T {
		i := num.RandomWithinRange(0, len(d.data)-1)
		return d.data[i]
	}
	if len(yieldFn) > 0 {
		yield = yieldFn[0]
	}

	d.yield = yield
	return d
}

// Yield calls an anonymous method that returns a single instance of T 'count' times, appending the results to the underlying data slice.
func (d *Data[T]) Yield(count uint) []T {
	out := make([]T, count)
	for i := uint(0); i < count; i++ {
		out[i] = d.yield()
		d.Append(out[i])
	}
	return out
}

func (d *Data[T]) Append(value ...T) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.data = append(d.data, value...)
}

func (d *Data[T]) Prepend(value ...T) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.data = append(value, d.data...)
}

func (d *Data[T]) Insert(index uint, value T) error {
	if index >= uint(len(d.data)) {
		return d.errorOutOfBounds(index)
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.data = append(d.data[:index], append([]T{value}, d.data[index:]...)...)
	return nil
}

func (d *Data[T]) Remove(index uint) error {
	if index >= uint(len(d.data)) {
		return d.errorOutOfBounds(index)
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.data = append(d.data[:index], d.data[index+1:]...)
	return nil
}

func (d *Data[T]) Select(low uint, high ...uint) ([]T, error) {
	if low >= uint(len(d.data)) {
		return nil, d.errorOutOfBounds(low)
	}

	if len(high) > 0 {
		// Selecting a range
		h := high[0]
		if h >= uint(len(d.data)) {
			return nil, d.errorOutOfBounds(h)
		} else if h < low {
			return nil, fmt.Errorf("high index %d must be greater than low index %d", h, low)
		}
		return d.data[low:h], nil
	}

	// Selecting a single index
	return []T{d.data[low]}, nil
}

func (d *Data[T]) SelectAll() []T {
	return d.data
}

func (d *Data[T]) errorOutOfBounds(index uint) error {
	return fmt.Errorf("index %d out of bounds [0,%d]", index, len(d.data)-1)
}
