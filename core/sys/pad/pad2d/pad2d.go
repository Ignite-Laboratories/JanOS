package pad2d

import (
	"github.com/ignite-laboratories/core/enum/direction/ordinal"
	"github.com/ignite-laboratories/core/sys/pad"
	"github.com/ignite-laboratories/core/sys/pad/scheme"
)

// UsingPattern pads the provided 2D slice using a 1D pattern.
func UsingPattern[T any, TSlice []T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() T, rolling ...bool) []TSlice {
	for d2 := range source {
		source[d2] = pad.UsingPattern(patternScheme, side, totalWidth, source[d2], pattern, rolling...)
	}
	return source
}

// Using2DPattern pads the provided 2D slice using a 2D pattern.
func Using2DPattern[T any, TSlice []T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() TSlice, rolling ...bool) []TSlice {
	return pad.UsingPattern(patternScheme, side, totalWidth, source, pattern, rolling...)
}

// Align pads every dimension to the same 'totalWidth' using the provided pattern.
func Align[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source [][]T, pattern func() T, rolling ...bool) [][]T {
	return Using2DPattern(patternScheme, side, totalWidth, source, func() []T {
		return pad.UsingPattern(patternScheme, side, totalWidth, []T{}, pattern, rolling...)
	}, rolling...)
}
