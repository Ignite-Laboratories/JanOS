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

// Precision is the global placeholder precision for floating point arithmetic.
var Precision uint = 256

// SeedRefractoryPeriod is the default amount of time a seed pool will retain its current random value set for.
// This allows a small batch of fixed random numbers to be referenced ad-hoc without neurons having to track their own concept of temporality.
var SeedRefractoryPeriod = 180 * time.Second

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
