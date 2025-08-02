package std

import (
	"fmt"
	"github.com/ignite-laboratories/core/math"
)

// XY is a general structure for holding generic (x,y) coordinate values.
type XY[T math.Numeric] struct {
	X T
	Y T
}

func (c XY[T]) String() string {
	return fmt.Sprintf("(%v, %v)", c.X, c.Y)
}
