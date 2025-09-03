package num

// A Breach indicates the amount that a bounded value overflowed or underflowed its boundaries.  The value is provided
// as an empty string when no breach occurred, or a string representation of a signed value.
type Breach string
