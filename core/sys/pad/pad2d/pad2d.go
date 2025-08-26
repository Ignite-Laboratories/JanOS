package pad2d

// UsingPattern pads the provided 2D slice using a 1D pattern.
//func UsingPattern[T any, TSlice []T](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []TSlice, pattern func(std.XY[uint]) T, rolling ...bool) []TSlice {
//	for d2 := range source {
//		source[d2] = pad.UsingPattern(patternScheme, side, totalWidth, source[d2], func(x uint) T {
//			coords := xy.From()
//			xy2 := std.XY[uint]{}
//			xy2.Set(x, uint(d2))
//			fmt.Println(x2y)
//			return pattern(xy2)
//		})
//	}
//	return source
//}

// Using2DPattern pads the provided 2D slice using a 2D pattern.
//func Using2DPattern[T any](patternScheme scheme.Scheme, side ordinal.Direction, totalWidth uint, source []T, pattern func(std.XY[uint]) T, rolling ...bool) []T {
//	return pad.UsingPattern(patternScheme, side, totalWidth, source, pattern, rolling...)
//}
