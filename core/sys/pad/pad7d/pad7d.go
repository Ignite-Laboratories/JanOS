package pad7d

//// UsingPattern pads the provided 7D slice using a 1D pattern.
//func UsingPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() T, rolling ...bool) []TSlice {
//	for d7 := range source {
//		for d6 := range source[d7] {
//			for d5 := range source[d7][d6] {
//				for d4 := range source[d7][d6][d5] {
//					for d3 := range source[d7][d6][d5][d4] {
//						for d2 := range source[d7][d6][d5][d4][d3] {
//							source[d7][d6][d5][d4][d3][d2] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6][d5][d4][d3][d2], pattern, rolling...)
//						}
//					}
//				}
//			}
//		}
//	}
//	return source
//}
//
//// Using2DPattern pads the provided 7D slice using a 2D pattern.
//func Using2DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() []T, rolling ...bool) []TSlice {
//	for d7 := range source {
//		for d6 := range source[d7] {
//			for d5 := range source[d7][d6] {
//				for d4 := range source[d7][d6][d5] {
//					for d3 := range source[d7][d6][d5][d4] {
//						source[d7][d6][d5][d4][d3] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6][d5][d4][d3], pattern, rolling...)
//					}
//				}
//			}
//		}
//	}
//	return source
//}
//
//// Using3DPattern pads the provided 7D slice using a 3D pattern.
//func Using3DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() [][]T, rolling ...bool) []TSlice {
//	for d7 := range source {
//		for d6 := range source[d7] {
//			for d5 := range source[d7][d6] {
//				for d4 := range source[d7][d6][d5] {
//					source[d7][d6][d5][d4] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6][d5][d4], pattern, rolling...)
//				}
//			}
//		}
//	}
//	return source
//}
//
//// Using4DPattern pads the provided 7D slice using a 4D pattern.
//func Using4DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() [][][]T, rolling ...bool) []TSlice {
//	for d7 := range source {
//		for d6 := range source[d7] {
//			for d5 := range source[d7][d6] {
//				source[d7][d6][d5] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6][d5], pattern, rolling...)
//			}
//		}
//	}
//	return source
//}
//
//// Using5DPattern pads the provided 7D slice using a 5D pattern.
//func Using5DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() [][][][]T, rolling ...bool) []TSlice {
//	for d7 := range source {
//		for d6 := range source[d7] {
//			source[d7][d6] = pad.UsingPattern(patternScheme, side, totalWidth, source[d7][d6], pattern, rolling...)
//		}
//	}
//	return source
//}
//
//// Using6DPattern pads the provided 7D slice using a 6D pattern.
//func Using6DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() [][][][][]T, rolling ...bool) []TSlice {
//	for i7 := range source {
//		source[i7] = pad.UsingPattern(patternScheme, side, totalWidth, source[i7], pattern, rolling...)
//	}
//	return source
//}
//
//// Using7DPattern pads the provided 7D slice using a 7D pattern.
//func Using7DPattern[T any, TSlice [][][][][][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() TSlice, rolling ...bool) []TSlice {
//	return pad.UsingPattern(patternScheme, side, totalWidth, source, pattern, rolling...)
//}
//
//func SeedAsymmetric[T any](patternScheme scheme.Scheme, side ordinal.Direction, xWidth, yWidth, zWidth, wWidth, aWidth, bWidth, cWidth uint, pattern func(x, y, z, w, a, b, c uint) T, rolling ...bool) [][][][][][][]T {
//	source := make([][][][][][][]T, 0)
//	c := uint(0)
//	return Using7DPattern(patternScheme, side, cWidth, source, func() [][][][][][]T {
//		b := uint(0)
//		outB := pad6d.Using6DPattern(patternScheme, side, bWidth, [][][][][][]T{}, func() [][][][][]T {
//			a := uint(0)
//			outA := pad5d.Using5DPattern(patternScheme, side, aWidth, [][][][][]T{}, func() [][][][]T {
//				w := uint(0)
//				outW := pad4d.Using4DPattern(patternScheme, side, wWidth, [][][][]T{}, func() [][][]T {
//					z := uint(0)
//					outZ := pad3d.Using3DPattern(patternScheme, side, zWidth, [][][]T{}, func() [][]T {
//						y := uint(0)
//						outY := pad2d.Using2DPattern(patternScheme, side, yWidth, [][]T{}, func() []T {
//							x := uint(0)
//							outX := pad.UsingPattern(patternScheme, side, xWidth, []T{}, func() T {
//								out := pattern(x, y, z, w, a, b, c)
//								x++
//								return out
//							}, rolling...)
//							y++
//							return outX
//						}, rolling...)
//						z++
//						return outY
//					}, rolling...)
//					w++
//					return outZ
//				}, rolling...)
//				a++
//				return outW
//			}, rolling...)
//			b++
//			return outA
//		}, rolling...)
//		c++
//		return outB
//	}, rolling...)
//}
//
//// Seed creates a new 7D[T] structure by calling the provided pattern function with coordinate
//func Seed[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, pattern func(x, y, z, w, a, b, c uint) T, rolling ...bool) [][][][][][][]T {
//	return SeedAsymmetric(patternScheme, side, totalWidth, totalWidth, totalWidth, totalWidth, totalWidth, totalWidth, totalWidth, pattern, rolling...)
//}
