package engine

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/name"
	"math"
	"time"
)

// New creates and configures a new instance of a neural impulse engine.
//
// You may optionally provide a name whilst creating your engine.
func New(named ...name.Given) *std.Engine {
	e := std.Engine{}
	e.Entity = std.NewEntity[name.Default](named...)
	e.MaxFrequency = math.MaxFloat64

	// Set up impulse regulation
	regulator := func(ctx std.Context) {
		if e.Resistance > 0 {
			time.Sleep(e.Resistance)
		}
	}

	e.Block(regulator, func(ctx std.Context) bool { return true }, false)
	return &e
}
