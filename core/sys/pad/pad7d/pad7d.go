package pad7d

import (
	"github.com/ignite-laboratories/core/enum/direction/ordinal"
	"github.com/ignite-laboratories/core/sys/pad"
	"github.com/ignite-laboratories/core/sys/pad/pad2d"
	"github.com/ignite-laboratories/core/sys/pad/pad3d"
	"github.com/ignite-laboratories/core/sys/pad/pad4d"
	"github.com/ignite-laboratories/core/sys/pad/pad5d"
	"github.com/ignite-laboratories/core/sys/pad/pad6d"
	"github.com/ignite-laboratories/core/sys/pad/scheme"
)

// UsingPattern pads the provided 7D slice using a 1D pattern.
func UsingPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() T, rolling ...bool) []TSlice {
	for d7 := range source {
		for d6 := range source[d7] {
			for d5 := range source[d7][d6] {
				for d4 := range source[d7][d6][d5] {
					for d3 := range source[d7][d6][d5][d4] {
						for d2 := range source[d7][d6][d5][d4][d3] {
							source[d7][d6][d5][d4][d3][d2] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6][d5][d4][d3][d2], pattern, rolling...)
						}
					}
				}
			}
		}
	}
	return source
}

// Using2DPattern pads the provided 7D slice using a 2D pattern.
func Using2DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() []T, rolling ...bool) []TSlice {
	for d7 := range source {
		for d6 := range source[d7] {
			for d5 := range source[d7][d6] {
				for d4 := range source[d7][d6][d5] {
					for d3 := range source[d7][d6][d5][d4] {
						source[d7][d6][d5][d4][d3] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6][d5][d4][d3], pattern, rolling...)
					}
				}
			}
		}
	}
	return source
}

// Using3DPattern pads the provided 7D slice using a 3D pattern.
func Using3DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() [][]T, rolling ...bool) []TSlice {
	for d7 := range source {
		for d6 := range source[d7] {
			for d5 := range source[d7][d6] {
				for d4 := range source[d7][d6][d5] {
					source[d7][d6][d5][d4] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6][d5][d4], pattern, rolling...)
				}
			}
		}
	}
	return source
}

// Using4DPattern pads the provided 7D slice using a 4D pattern.
func Using4DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() [][][]T, rolling ...bool) []TSlice {
	for d7 := range source {
		for d6 := range source[d7] {
			for d5 := range source[d7][d6] {
				source[d7][d6][d5] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6][d5], pattern, rolling...)
			}
		}
	}
	return source
}

// Using5DPattern pads the provided 7D slice using a 5D pattern.
func Using5DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() [][][][]T, rolling ...bool) []TSlice {
	for d7 := range source {
		for d6 := range source[d7] {
			source[d7][d6] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6], pattern, rolling...)
		}
	}
	return source
}

// Using6DPattern pads the provided 7D slice using a 6D pattern.
func Using6DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() [][][][][]T, rolling ...bool) []TSlice {
	for i7 := range source {
		source[i7] = pad.UsingPattern(patternScheme, side, totalWidth, source[i7], pattern, rolling...)
	}
	return source
}

// Using7DPattern pads the provided 7D slice using a 7D pattern.
func Using7DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() TSlice, rolling ...bool) []TSlice {
	return pad.UsingPattern(patternScheme, side, totalWidth, source, pattern, rolling...)
}

// Align pads every dimension to the same 'totalWidth' using the provided pattern.
func Align[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source [][][][][][][]T, pattern func() T, rolling ...bool) [][][][][][][]T {
	return Using7DPattern(patternScheme, side, totalWidth, source, func() [][][][][][]T {
		return pad6d.Using6DPattern(patternScheme, side, totalWidth, [][][][][][]T{}, func() [][][][][]T {
			return pad5d.Using5DPattern(patternScheme, side, totalWidth, [][][][][]T{}, func() [][][][]T {
				return pad4d.Using4DPattern(patternScheme, side, totalWidth, [][][][]T{}, func() [][][]T {
					return pad3d.Using3DPattern(patternScheme, side, totalWidth, [][][]T{}, func() [][]T {
						return pad2d.Using2DPattern(patternScheme, side, totalWidth, [][]T{}, func() []T {
							return pad.UsingPattern(patternScheme, side, totalWidth, []T{}, pattern, rolling...)
						}, rolling...)
					}, rolling...)
				}, rolling...)
			}, rolling...)
		}, rolling...)
	}, rolling...)
}
