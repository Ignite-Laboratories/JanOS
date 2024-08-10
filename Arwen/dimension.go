package Arwen

import (
	"github.com/Ignite-Laboratories/JanOS/Logic/Math"
)

// Dimension represents an instantaneous reading of a value in time.
type Dimension struct {
	Value  float64
	Name   string
	Symbol Math.Symbol
}

func NewDimension(name string, symbol Math.Symbol, defaultValue ...float64) *Dimension {
	d := &Dimension{
		Name:   name,
		Symbol: symbol,
	}
	if defaultValue != nil {
		d.Value = defaultValue[0]
	}
	return d
}
