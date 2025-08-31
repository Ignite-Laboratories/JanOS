package std

import (
	"core/sys/name/format"
	"core/sys/num"
)

type Cursor[Data any, T num.Primitive] struct {
	_cursor[T]
	data []Data

	initialized bool
}

func (c *Cursor[Data, T]) sanityCheck() {
	minI := T(0)
	maxI := T(len(c.data) - 1)
	if len(c.data) == 0 {
		maxI = 0
	}

	if !c.initialized {
		_c := _cursor[T]{}
		_c.Entity = NewEntity[format.Default]()
		_c.SetBoundaries(minI, maxI)
		_c.Set(0)
		c._cursor = _c
		c.initialized = true
	} else {
		_ = c.SetBoundaries(minI, maxI)
		_ = c.Set(c.I.Value())
	}
}
