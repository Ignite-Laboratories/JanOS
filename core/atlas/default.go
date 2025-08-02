// Package atlas provides the globally referencable values used by JanOS.
package atlas

import (
	"github.com/ignite-laboratories/core/std"
	"time"
)

// ObservanceWindow is the default dimensional window of observance - 2 seconds.
var ObservanceWindow = 2 * time.Second

// TrimFrequency sets the default global frequency for dimensional trimmers.
var TrimFrequency = 1024.0 //hz

// Impulse is the global impulse engine.
var Impulse *std.Engine = std.NewEngine()

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
