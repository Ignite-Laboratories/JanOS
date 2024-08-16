package JanOS

import (
	"time"
)

// Symbol represents a mathematical symbol.
type Symbol string

// Dimension represents a mathematical value and its associated historical timeline.
type Dimension struct {
	Name     string
	Symbol   Symbol
	Timeline *timeline
}

type dimensionManager struct {
	dimensions map[string]*Dimension
}

func newDimensionManager() *dimensionManager {
	return &dimensionManager{
		dimensions: make(map[string]*Dimension),
	}
}

// GetNamedValue returns the assigned name to this instance.
func (mgr *dimensionManager) GetNamedValue() string {
	return "Dimensions"
}

// GetDimension references a previously stored dimension.
func (mgr *dimensionManager) GetDimension(name string) *Dimension {
	return mgr.dimensions[name]
}

// SetValue seeks to the appropriate position in time and replaces the values on the remainder of the buffer.
// The mentality is that when you set a value in time it will hold that value until it is changed.  We take
// time for granted in our environment, meaning that the setting of a value should always assume the future
// will hold that value ad infinitum.
func (d *Dimension) SetValue(instant time.Time, value float64) {
	Universe.Printf(Universe.Dimensions, "Set '%s' [%s] = %f", d.Name, string(d.Symbol), value)
	d.Timeline.setValue(instant, value)
}

// GetValue seeks to the appropriate position in time and gets the value on the buffer at that instant.
func (d *Dimension) GetValue(instant time.Time) float64 {
	return d.Timeline.GetInstant(instant)
}

// NewDimension creates a new dimension and sets its timeline to the provided default value.
func (mgr *dimensionManager) NewDimension(name string, symbol Symbol, defaultValue float64) *Dimension {
	Universe.Printf(mgr, "Let '%s' [%s] = %f", name, string(symbol), defaultValue)
	d := &Dimension{
		Name:     name,
		Symbol:   symbol,
		Timeline: mgr.newTimeline(name, symbol, defaultValue, Universe.BufferLength),
	}
	mgr.dimensions[name] = d
	return d
}
