package std

import (
	"core/sys/name/format"
)

type Cursor[Data any] struct {
	index _cursor[uint]
	data  []Data

	initialized bool
}

func (c *Cursor[Data]) SanityCheck() {
	minI := uint(0)
	maxI := uint(len(c.data)) - 1
	if len(c.data) == 0 {
		maxI = 0
	}

	if !c.initialized {
		_c := _cursor[uint]{}
		_c.Entity = NewEntity[format.Default]()
		_c.SetBoundaries(minI, maxI)
		_c.Set(0)
		c.index = _c
		c.initialized = true
	} else {
		_ = c.index.SetBoundaries(minI, maxI)
		_ = c.index.Set(c.index.I.Value())
	}
}
