package support

import "reflect"

// Deduplicate removes duplicate entries from the provided data.
func Deduplicate[T any](data []T) []T {
	seen := make(map[any]struct{}, len(data))
	unique := make([]T, 0, len(data))
	for _, v := range data {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		unique = append(unique, v)
	}
	return unique
}

// IsComparable performs a runtime check to determine if the provided object is a 'comparable' type.
func IsComparable(v any) bool {
	t := reflect.TypeOf(v)
	if t == nil { // nil interface value
		return true // comparing nil interface values is well-defined
	}
	return t.Comparable()
}

// AllSameTypes returns true if all the provided instances are of the same type.
func AllSameTypes(vals ...any) bool {
	if len(vals) <= 1 {
		return true
	}
	t0 := reflect.TypeOf(vals[0])
	for _, v := range vals[1:] {
		if reflect.TypeOf(v) != t0 {
			return false
		}
	}
	return true
}

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
