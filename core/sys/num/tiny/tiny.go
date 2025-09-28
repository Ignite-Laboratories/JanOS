package tiny

import (
	"reflect"

	"git.ignitelabs.net/janos/core/sys/atlas"
	"git.ignitelabs.net/janos/core/sys/num"
)

func filterBounds(operands ...any) (num.Bounds, []any) {
	// NOTE: All fields are dereferenced locally to ensure the calculation is stateful in-flight.

	dereference := func(v any) any {
		rv := reflect.ValueOf(v)
		for rv.Kind() == reflect.Ptr {
			if rv.IsNil() {
				return nil
			}
			rv = rv.Elem()
		}
		return rv.Interface()
	}

	base := atlas.Base
	bounds := num.Bounds{
		Base: &base,
	}
	out := make([]any, 0, len(operands))
	for _, operand := range operands {
		if op, ok := operand.(num.Bounds); ok {
			if op.Minimum != nil {
				bounds.Minimum = dereference(op.Minimum)
			}
			if op.Maximum != nil {
				bounds.Maximum = dereference(op.Maximum)
			}
			if op.Clamp != nil {
				clamp := *op.Clamp
				bounds.Clamp = &clamp
			}
			if op.Base != nil {
				b := *op.Base
				bounds.Base = &b
			}
		} else {
			out = append(out, operand)
		}
	}
	return bounds, out
}

func Compare(a, b any) int {
	// TODO: Handle the ∞ character as 'infinity' specially here
	panic("comparison operations are not yet implemented")
}

func Add(operands ...any) any {
	//bounds, out := filterBounds(operands...)
	panic("addition operations are not yet implemented")

	// If the operand is a []byte, the first index should indicate it's base - and it should panic if this is missing!
	// If the operand is a string, its base should be indicated with a # character at the beginning (or 10 if omitted)

	//  ~123.‾45 <- a plain base 10 number
	//  10#~123.‾45 <- an annotated base 10 number
	//  2#1010.‾010 <- a base 2 number
	//  17#(~ AA BB . ‾ 10) <- a base 17+ number
	//  256#(AA BB F0) <- A base 256 number
	//
	//  { 10, 4, 2} <- The base 10 natural number '42' in []byte form
	//  { 2, 1, 0, 1, 1 } <- A base 2 natural number in []byte form
	//  { 256, AA, BB, F0 } <- A base 256 natural number in []byte form
}

func Subtract(operands ...any) any {
	panic("subtraction operations are not yet implemented")
}

func Multiply(operands ...any) any {
	panic("multiplication operations are not yet implemented")
}

func Divide(operands ...any) any {
	panic("division operations are not yet implemented")
}

func Modulo(operands ...any) any {
	panic("modulo operations are not yet implemented")
}

func Floor(operands ...any) any {
	panic("flooring operations are not yet implemented")
}

func Ceiling(operands ...any) any {
	panic("ceiling operations are not yet implemented")
}

// ParseAs takes in any operand and evaluates it into the provided type.  If provided a string formula, this will
// execute it and yield the desired output.  You may optionally provide a num.Bounds defining the boundaries
// of the evaluated result.
func ParseAs[TOut any](operand any, bounds ...num.Bounds) TOut {
	panic("parsing operations are not yet implemented")
}
