package movement

//NOTE NOTE NOTE: Set the boundaries at the beginning of during every emit call BY CONVENTION since Go's slices are "ragged"

//func Function1D[T any](cursor *std.Cursor[uint], data ...T) (x std.Emit[T]) {
//	xSingleR := func(i int) T {
//		var out T
//		_ = cursor.SetBoundaries(0, uint(len(data)-1))
//		if i < 0 {
//			_ = cursor.AddOrSubtract(i)
//			out = data[cursor.Value()]
//		} else {
//			out = data[cursor.Value()]
//			_ = cursor.AddOrSubtract(i)
//		}
//		return out
//	}
//	xManyR := func(i int) []T {}
//	xSingleA := func(i int) T {
//		_ = cursor.SetBoundaries(0, uint(len(data)-1))
//		_ = cursor.Set(uint(i))
//		return data[cursor.Value()]
//	}
//	xManyA := func(i int) []T {}
//
//	return std.NewEmit[T](xSingleR, xManyR, xSingleA, xManyA)
//}
//
//func Function2D[T any](cursor *std.Cursor[uint], data ...[]T) (x std.Emit[T], y std.Emit[[]T]) {
//	xSingleR := func(i int) T {
//		var out T
//		_ = cursor.SetBoundaries(0, uint(len(data)-1))
//		if i < 0 {
//			_ = cursor.AddOrSubtract(i)
//			out = data[][cursor.Value()]
//		} else {
//			out = data[][cursor.Value()]
//			_ = cursor.AddOrSubtract(i)
//		}
//		return out
//	}
//	return x, y
//}
//
//func Function3D[T any](data ...[][]T) (x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T]) {
//	x, y, z, _, _, _, _ = function[T](3, data)
//	return x, y, z
//}
//
//func Function4D[T any](data ...[][][]T) (x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], w std.Emit[[][][]T]) {
//	x, y, z, w, _, _, _ = function[T](4, data)
//	return x, y, z, w
//}
//
//func Function5D[T any](data ...[][][][]T) (x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], w std.Emit[[][][]T], a std.Emit[[][][][]T]) {
//	x, y, z, w, a, _, _ = function[T](5, data)
//	return x, y, z, w, a
//}
//
//func Function6D[T any](data ...[][][][][]T) (x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], w std.Emit[[][][]T], a std.Emit[[][][][]T], b std.Emit[[][][][][]T]) {
//	x, y, z, w, a, b, _ = function[T](6, data)
//	return x, y, z, w, a, b
//}
//
//func Function7D[T any](data ...[][][][][][]T) (x std.Emit[T], y std.Emit[[]T], z std.Emit[[][]T], w std.Emit[[][][]T], a std.Emit[[][][][]T], b std.Emit[[][][][][]T], c std.Emit[[][][][][][]T]) {
//	return function[T](7, data)
//}
//
//// From creates a new std.Pattern which can infinitely walk through the provided data either westbound or eastbound.
////
//// NOTE: This will create a single element 'zero' instance pattern of T if provided no data.
//func From[T any](data ...T) std.Emit[T] {
//	if len(data) == 0 {
//		var zero T
//		data = append(data, zero)
//	}
//
//	c, _ := bounded.By[uint](0, 0, uint(len(data)-1))
//	walkEast := func(i uint) T {
//		out := data[c.Value()]
//		_ = c.Increment(i)
//		return out
//	}
//	walkWest := func(i uint) T {
//		_ = c.DecrementPtr(i)
//		return data[c.Value()]
//	}
//	walkTo := func(i uint) T {
//		_ = c.SetPtr(i)
//		return data[c.Value()]
//	}
//	yieldEast := func(i uint) []T {
//		out := make([]T, i)
//		for j := uint(0); j < i; j++ {
//			out[j] = walkEast(1)
//		}
//		return out
//	}
//	yieldWest := func(i uint) []T {
//		out := make([]T, i)
//		for j := uint(0); j < i; j++ {
//			out[j] = walkWest(1)
//		}
//		return out
//	}
//	yieldTo := func(i uint) []T {
//		if i > c.Value() {
//			return yieldEast(i - c.Value())
//		} else if i < c.Value() {
//			return yieldWest(c.Value() - i)
//		}
//		return []T{data[c.Value()]}
//	}
//
//	return internal.NewPattern[T](&c, walkEast, walkWest, walkTo, yieldEast, yieldWest, yieldTo, data...)
//}
