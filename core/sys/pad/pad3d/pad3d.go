package pad3d

//// UsingPattern pads the provided 3D slice using a 1D pattern.
//func UsingPattern[T any, TSlice [][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() T, rolling ...bool) []TSlice {
//	for d3 := range source {
//		for d2 := range source[d3] {
//			source[d3][d2] = pad.UsingPattern(patternScheme, side, totalWidth, source[d3][d2], pattern, rolling...)
//		}
//	}
//	return source
//}
//
//// Using2DPattern pads the provided 3D slice using a 2D pattern.
//func Using2DPattern[T any, TSlice [][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() []T, rolling ...bool) []TSlice {
//	for d3 := range source {
//		source[d3] = pad.UsingPattern(patternScheme, side, totalWidth, source[d3], pattern, rolling...)
//	}
//	return source
//}
//
//// Using3DPattern pads the provided 3D slice using a 3D pattern.
//func Using3DPattern[T any, TSlice [][]T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func() TSlice, rolling ...bool) []TSlice {
//	return pad.UsingPattern(patternScheme, side, totalWidth, source, pattern, rolling...)
//}
//
//// Align pads every dimension to the same 'totalWidth' using the provided pattern.
//func Align[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source [][][]T, pattern func() T, rolling ...bool) [][][]T {
//	return Using3DPattern(patternScheme, side, totalWidth, source, func() [][]T {
//		return pad2d.Using2DPattern(patternScheme, side, totalWidth, [][]T{}, func() []T {
//			return pad.UsingPattern(patternScheme, side, totalWidth, []T{}, pattern, rolling...)
//		}, rolling...)
//	}, rolling...)
//}
