package support

import (
	"math/rand/v2"
	"reflect"
	"slices"
	"strings"
	"unicode"
)

// SetCase changes the case of the input string.  If no indices are provided, the entire string is affected.
// Otherwise, only the provided indices are manipulated.
func SetCase(input string, upper bool, indices ...int) string {
	if len(indices) == 0 {
		if upper {
			return strings.ToUpper(input)
		}
		return strings.ToLower(input)
	}

	runes := []rune(input)
	for _, ii := range indices {
		if upper {
			runes[ii] = unicode.ToUpper(runes[ii])
		} else {
			runes[ii] = unicode.ToLower(runes[ii])
		}
	}
	return string(runes)
}

// ShuffleSet clones and then shuffles the provided set using 'slices' and 'math/rand/v2', respectively.
func ShuffleSet[T any](set []T) []T {
	s := slices.Clone(set)
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}

// IsAlphaNumeric returns whether the provided string contains only the alphanumeric characters [0-9, a-z, A-Z]
func IsAlphaNumeric(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

// Deduplicate removes duplicate entries from the provided data.
func Deduplicate[T comparable](data []T) []T {
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
