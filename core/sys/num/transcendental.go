package num

import "core/enum/transcendental"

func init() {

}

var _piMap map[uint]Realized
var _eMap map[uint]Realized
var _sqrt2Map map[uint]Realized

// IsTranscendental returns which transcendental constant the input Realized is.
//
// See transcendental.Transcendental, Pi, E, and Sqrt2
func IsTranscendental(r Realized) transcendental.Transcendental {
	// TODO: implement this
	return transcendental.Non
}

// Pi represents the transcendental constant 'π' in your requested base.
//
// NOTE: If no placeholder value is provided, this will use atlas.Precision.
func Pi(base uint16, placeholders ...uint) Realized {
	// TODO: implement this
	panic("transcendental calculation is not yet implemented")
}

// E represents the transcendental constant of Euler's number in your requested base.
//
// NOTE: If no placeholder value is provided, this will use atlas.Precision.
func E(base uint16, placeholders ...uint) Realized {
	// TODO: implement this
	panic("transcendental calculation is not yet implemented")
}

// Sqrt2 represents the transcendental constant of √2 in your requested base.
//
// NOTE: If no placeholder value is provided, this will use atlas.Precision.
func Sqrt2(base uint16, placeholders ...uint) Realized {
	// TODO: implement this
	panic("transcendental calculation is not yet implemented")
}
