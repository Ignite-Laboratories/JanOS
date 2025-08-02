package tiny

import (
	"github.com/ignite-laboratories/core/std"
	complex2 "github.com/ignite-laboratories/core/sys/tiny/complex"
	"github.com/ignite-laboratories/core/sys/tiny/index"
	"github.com/ignite-laboratories/core/sys/tiny/natural"
	real2 "github.com/ignite-laboratories/core/sys/tiny/real"
)

// Operable represents the basic logically operable types.
//
// See Bit, Measurement, Phrase, Natural, Real, Complex, and Index
type Operable interface {
	std.Bit | byte | std.Measurement[any] | std.Phrase[any] | natural.Natural | real2.Real | complex2.Complex | index.Index
}
