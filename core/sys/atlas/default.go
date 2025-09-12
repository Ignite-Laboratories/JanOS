package atlas

import (
	"time"
)

// PrintPreamble indicates if JanOS should print it's preamble.
var PrintPreamble = true

// ObservanceWindow is the default dimensional window of observance - 2 seconds.
var ObservanceWindow = 2 * time.Second

// TrimFrequency sets the default global frequency for dimensional trimmers.
var TrimFrequency = 1024.0 //hz

// Precision is the global maximum number of placeholders to consider in the fractional part of a num.Realized.
var Precision uint = 256

// PrecisionMinimum indicates the minimum width to synthesize irrational or periodic fractional components to during
// printing. 𝑡𝑖𝑛𝑦 will also round the value accordingly, meaning √2 would render as "~1.4142136" instead of "~1.4142135"
//
// NOTE: This defaults to a 7 placeholder minimum.
var PrecisionMinimum uint = 7

// Base defines the global default base for all calculation.
//
// NOTE: This defaults to base₁₀.
var Base uint16 = 10

// SeedRefractoryPeriod is the default amount of time a seed pool will retain its current random value set for.
// This allows a small batch of fixed random numbers to be referenced ad-hoc without neurons having to track their own concept of temporality.
var SeedRefractoryPeriod = 180 * time.Second

// IncludeNilBits indicates if a num.Bit value of '219' is considered to be acceptable as a 'nil' value.
var IncludeNilBits = false

// CompactVectors indicates if vector string functions should print in "compact" or "full" form.
//
// Full:
//
//	xy[uint]{xVal, yVal}(givenName)
//
// Compact:
//
//	{x: xVal, y: yVal}
var CompactVectors = false

/**
Constant References
*/

// True is a constantly referenceable true - please don't change it!
//
// See False
var True bool = true

// False is a constantly referenceable false - please don't change it!
//
// See True
var False bool = false
