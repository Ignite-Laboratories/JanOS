// Go
package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Accessor struct {
	IsPointer bool
	Get       func() any
	Set       func(any) error // no-op or error for non-pointers
}

// Go
func captureAll(ops ...any) []Accessor {
	out := make([]Accessor, len(ops))
	for i, op := range ops {
		out[i] = makeAccessor(op)
	}
	return out
}

func makeAccessor(x any) Accessor {
	v := reflect.ValueOf(x)
	if !v.IsValid() {
		return Accessor{
			IsPointer: false,
			Get:       func() any { return nil },
			Set:       func(any) error { return errors.New("invalid value") },
		}
	}

	if v.Kind() == reflect.Ptr {
		return Accessor{
			IsPointer: true,
			Get: func() any {
				if v.IsNil() {
					return nil
				}
				return v.Elem().Interface()
			},
			Set: func(val any) error {
				if v.IsNil() {
					return errors.New("nil pointer")
				}
				dst := v.Elem()
				src := reflect.ValueOf(val)
				if !src.Type().AssignableTo(dst.Type()) {
					return fmt.Errorf("cannot assign %v to %v", src.Type(), dst.Type())
				}
				if !dst.CanSet() {
					return errors.New("destination not settable")
				}
				dst.Set(src)
				return nil
			},
		}
	}

	// Not a pointer
	return Accessor{
		IsPointer: false,
		Get:       func() any { return x },
		Set:       func(any) error { return errors.New("not a pointer") },
	}
}
