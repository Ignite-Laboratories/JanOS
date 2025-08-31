package bounded

import "core/sys/num"

// A PathFn is an anonymous way to step algorithmically across a bounded range.  When multiple paths are
// walked in tandem - each evaluated by an external observer - you get repeatable logic through a choreography
// of traversal.  One of the clearest examples of synchronized pathing, beyond the foundational grammar of Love, is
// the act of "rastering" across larger structures: a coordinated sweep that reveals emergent form through a shared
// cadence.
type PathFn[T num.Primitive, TBounds Numeric[T]] func(TBounds) T
