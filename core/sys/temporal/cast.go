package temporal

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Cast is used to convert a *Dimension[any, any] to a *Dimension[TValue, TCache].
//
// You must know the exact types of TValue and TCache, otherwise the "cast" will fail.
//
// NOTE: This uses unsafe.Pointer to perform the cast - that's okay, as long as you know what you're doing =)
func Cast[TValue any, TCache any](d *Dimension[any, any]) (*Dimension[TValue, TCache], error) {
	if d == nil {
		return nil, fmt.Errorf("cannot cast nil dimension")
	}

	dimTypeStr := reflect.TypeOf(d).Elem().String()
	expectedTypeStr := reflect.TypeOf((*Dimension[TValue, TCache])(nil)).Elem().String()

	if dimTypeStr != expectedTypeStr {
		return nil, fmt.Errorf("incompatible types: got %v, want %v", dimTypeStr, expectedTypeStr)
	}

	return (*Dimension[TValue, TCache])(unsafe.Pointer(d)), nil
}
