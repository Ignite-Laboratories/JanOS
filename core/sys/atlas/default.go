package atlas

import (
	"time"

	"git.ignitelabs.net/janos/core/sys/rec"
)

// Verbose gets or sets the rec.Verbose value.  If no value is provided, it only gets - if a value is provided, it sets and gets.
func Verbose(set ...bool) bool {
	refresh()
	if len(set) > 0 {
		rec.Verbose = set[0]
	}
	return rec.Verbose
}

// Silent gets or sets the rec.Silent value.  If no value is provided, it only gets - if a value is provided, it sets and gets.
func Silent(set ...bool) bool {
	refresh()
	if len(set) > 0 {
		rec.Silent = set[0]
	}
	return rec.Silent
}

// PrintPreamble indicates if JanOS should print its preamble.
var PrintPreamble = true

// ShutdownTimeout is the default amount of time that JanOS will allow cleanup operations within during shutdown.
var ShutdownTimeout = 5 * time.Second

// Record holds the temporal history of this instance.
var Record []byte

// ObservanceWindow is the default dimensional window of observance - 2 seconds.
var ObservanceWindow = 5 * time.Second

// ObservedMinimum is the minimum number of elements a temporal buffer should hold, regardless of the observance window.
//
// NOTE: This can never be set to 0 - it will default back to 7.
var ObservedMinimum = uint(7)

// TrimFrequency sets the default global frequency for dimensional trimmers.
var TrimFrequency = 1024.0 //hz

// Precision is the global maximum number of placeholders to consider in the fractional part of a num.Realized.
var Precision uint = 256

// PrecisionMinimum indicates the minimum width to synthesize irrational or periodic fractional components to during
// printing. 𝑡𝑖𝑛𝑦 will also round the value accordingly, meaning √2 would render as "~1.4142136" instead of "~1.4142135"
//
// NOTE: This defaults to a 7 placeholder minimum.
var PrecisionMinimum uint = 7

// Radix defines the global default base for all calculation.
//
// NOTE: This defaults to base₁₀.
var Radix uint = 10

// SeedRefractoryPeriod is the default amount of time a seed pool will retain its current random value set for.
// This allows a small batch of fixed random numbers to be referenced ad-hoc without neuron having to track their own concept of temporality.
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

// SynapticChannelLimit defines the maximum number of signals a synapse channel can receive before blocking - defaulting to 2¹⁶
var SynapticChannelLimit = uint(1 << 16)
