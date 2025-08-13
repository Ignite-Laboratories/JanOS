package std

import (
	"github.com/ignite-laboratories/core/std/num"
)

// Cursor is a variant of Bounded[T] that operates entirely off of pointer receivers, rather than value receivers.
// The functionality is merely passed through to Bounded and all methods are maintained in parity with each other.
type Cursor[T num.Primitive] Bounded[T]

// NewCursor creates a new instance of Cursor[T].
//
// NOTE: While you can call this directly, the convention is to use the 'std/Cursor' package.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func NewCursor[T num.Primitive](value, minimum, maximum T, clamp ...bool) (Cursor[T], error) {
	c := len(clamp) > 0 && clamp[0]
	cur := Cursor[T]{
		value:   T(0),
		minimum: minimum,
		maximum: maximum,
		Clamp:   c,
	}
	err := cur.Set(value)
	return cur, err
}

// NewCursorDefault returns a standard *Cursor[uint] bound to [0,0].
func NewCursorDefault() *Cursor[uint] {
	c, _ := NewCursor[uint](0, 0, 0)
	return &c
}

func (bnd *Cursor[T]) ptrHelper(set Bounded[T]) {
	bnd.value = set.value
	bnd.minimum = set.minimum
	bnd.maximum = set.maximum
	bnd.Clamp = set.Clamp
}

// Value returns the currently held Cursor value.
func (bnd Cursor[T]) Value() T {
	return bnd.value
}

// Minimum returns the current minimum boundary.
func (bnd Cursor[T]) Minimum() T {
	return bnd.minimum
}

// Maximum returns the current maximum boundary.
func (bnd Cursor[T]) Maximum() T {
	return bnd.maximum
}

// Increment adds 1 or the provided count to the direct memory address of the bound value.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Cursor[T]) Increment(count ...T) error {
	set, err := Bounded[T](*bnd).Increment(count...)
	bnd.ptrHelper(set)
	return err
}

// Decrement subtracts 1 or the provided count from the bound value as a pointer function.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Cursor[T]) Decrement(count ...T) error {
	set, err := Bounded[T](*bnd).Decrement(count...)
	bnd.ptrHelper(set)
	return err
}

// SetAll sets the value and boundaries of a pointer to Cursor[T] all in one operation, preventing multiple calls to Set().
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Cursor[T]) SetAll(value, a, b T, clamp ...bool) error {
	set, err := Bounded[T](*bnd).SetAll(value, a, b, clamp...)
	bnd.ptrHelper(set)
	return err
}

// SetBoundariesFromType sets the boundaries to the implied limits of a pointer to the Cursor type before calling Set(current value).
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Cursor[T]) SetBoundariesFromType() error {
	set, err := Bounded[T](*bnd).SetBoundariesFromType()
	bnd.ptrHelper(set)
	return err
}

// SetBoundaries sets the boundaries of a pointer to Cursor before calling Set(current value).
//
// NOTE: The boundary parameters are evaluated to ensure the lower bound is always the 'minimum'
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Cursor[T]) SetBoundaries(a, b T) error {
	set, err := Bounded[T](*bnd).SetBoundaries(a, b)
	bnd.ptrHelper(set)
	return err
}

// Normalize converts the Cursor value to a float64 unit vector in the interval [0.0, 1.0],
// by linearly mapping the value from its Cursor interval's [minimum, maximum]. A value equal
// to minimum maps to 0.0, a value equal to maximum maps to 1.0, and values in between
// are linearly interpolated.
func (bnd *Cursor[T]) Normalize() float64 {
	return Bounded[T](*bnd).Normalize()
}

// Normalize32 converts the Cursor value to a float32 unit vector in the range [0.0, 1.0],
// where the Cursor minimum maps to 0.0 and the Cursor maximum maps to 1.0.
func (bnd *Cursor[T]) Normalize32() float32 {
	return float32(bnd.Normalize())
}

// SetFromNormalized sets the value of a pointer to Cursor using a float64 unit vector from the [0.0, 1.0]
// range, where 0.0 maps to the Cursor minimum and 1.0 maps to the Cursor maximum.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Cursor[T]) SetFromNormalized(normalized float64) error {
	set, err := Bounded[T](*bnd).SetFromNormalized(normalized)
	bnd.ptrHelper(set)
	return err
}

// SetFromNormalized32 sets the value of a pointer to Cursor using a float32 unit vector from the [0.0, 1.0]
// range, where 0.0 maps to the Cursor minimum and 1.0 maps to the Cursor maximum.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Cursor[T]) SetFromNormalized32(normalized float32) error {
	set, err := Bounded[T](*bnd).SetFromNormalized32(normalized)
	bnd.ptrHelper(set)
	return err
}

// Set sets the value of a pointer to a Cursor object and automatically handles when the value exceeds the boundaries.
//
// NOTE: This will return a safely ignorable 'under' or 'over' error if the value exceeded the boundaries.
func (bnd *Cursor[T]) Set(value T) error {
	set, err := Bounded[T](*bnd).Set(value)
	bnd.value = set.Value()
	return err
}
