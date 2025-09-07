package num

import (
	"core/enum/transcendental"
	"core/sys/atlas"
)

func init() {
	// Kick off a standard 'cached' value for each transcendental from bases 2-16 on initialization.
	for i := uint16(2); i < 17; i++ {
		Transcendental.Pi(i, &atlas.Precision)
		Transcendental.E(i, &atlas.Precision)
	}
}

type _transcendental struct {
	piMap map[uint]Realized
	eMap  map[uint]Realized
}

var Transcendental = _transcendental{
	piMap: make(map[uint]Realized),
	eMap:  make(map[uint]Realized),
}

// Is returns which transcendental constant the input Realized is.
//
// See transcendental.Number, Pi, E, and Sqrt2
func (t _transcendental) Is(r Realized) transcendental.Number {
	// NOTE: This will need to handle this as such:
	// If the realized is smaller than atlas.Precision, it's not transcendental
	// Otherwise, check to the realized's precision.
	// TODO: implement this
	return transcendental.Non
}

func (t _transcendental) From(base uint16, identifier transcendental.Number) Realized {
	// TODO: implement this
	// panic("transcendental calculation is not yet implemented")
	return Realized{}
}

// Pi represents the transcendental constant 'π' in your requested base.
//
// NOTE: If no placeholder value is provided, this will use atlas.Precision.
func (t _transcendental) Pi(base uint16, placeholders ...*uint) Realized {
	// TODO: implement this
	// panic("transcendental calculation is not yet implemented")
	return Realized{}
}

// E represents the transcendental constant of Euler's number in your requested base.
//
// NOTE: If no placeholder value is provided, this will use atlas.Precision.
func (t _transcendental) E(base uint16, placeholders ...*uint) Realized {
	// TODO: implement this
	// panic("transcendental calculation is not yet implemented")
	return Realized{}
}
