package xaxis

// Any represents any normalized movement space along the X-axis.
//
// See Any, Negative [-1], Static [0], and Positive [1]
type Any interface {
	Negative | Static | Positive
}

// Negative represents the negative normal along the X-axis.
//
// See Any, Negative [-1], Static [0], and Positive [1]
type Negative byte

// Static represents no movement along the X-axis.
//
// See Any, Negative [-1], Static [0], and Positive [1]
type Static byte

// Positive represents the positive normal along the X-axis
//
// See Any, Negative [-1], Static [0], and Positive [1]
type Positive byte
