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
