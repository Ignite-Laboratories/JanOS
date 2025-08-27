package pad2d

import (
	"core/enum/direction/ordinal"
	"core/std"
	"core/sys/pad"
	"core/sys/pad/scheme"
)

// UsingPattern pads the provided 2D slice using a 1D pattern.
func UsingPattern[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []T, pattern func(uint) T) []T {
	return pad.UsingPattern(patternScheme, side, totalWidth, source, pattern)
}

// Using2DPattern pads the provided 2D slice using a 2D pattern.
func Using2DPattern[T any, TSlice []T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func(std.XY[uint]) T) []TSlice {
	for y := range source {
		source[y] = pad.UsingPattern(patternScheme, side, totalWidth, source[y], func(x uint) T {
			return pattern(*std.NewXY[uint](x, uint(y)))
		})
	}
	return source
}
