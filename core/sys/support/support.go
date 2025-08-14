package support

import "reflect"

func SliceDepth(v any) (depth int, elem reflect.Type, ok bool) {
	t := reflect.TypeOf(v)
	if t == nil {
		return 0, nil, false
	}

	// Optional: unwrap pointers so you can pass in *[]T, **[]T, etc.
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		if t == nil {
			return 0, nil, false
		}
	}

	for t.Kind() == reflect.Slice {
		depth++
		t = t.Elem()
	}
	if depth == 0 {
		return 0, nil, false
	}
	return depth, t, true
}

// IsMultiDimensionalSlice is true if the value is at least [][]T (depth >= 2).
func IsMultiDimensionalSlice(v any) bool {
	d, _, ok := SliceDepth(v)
	return ok && d >= 2
}
