package tiny

import (
	"reflect"

	"git.ignitelabs.net/janos/core/sys/num"
)

func boundaryFilter(operands ...any) (num.Bounds, []any) {
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

	var bounds num.Bounds
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

func Add(operands ...any) func(...uint16) any {
	return func(...uint16) any {
		//bounds, out := boundaryFilter(operands...)
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
}

func Subtract(operands ...any) func(...uint) any {
	panic("subtraction operations are not yet implemented")
}

func Multiply(operands ...any) func(...uint) any {
	panic("multiplication operations are not yet implemented")
}

func Divide(operands ...any) func(...uint) any {
	panic("division operations are not yet implemented")
}

func Modulo(operands ...any) func(...uint) any {
	panic("modulo operations are not yet implemented")
}

func Floor(operands ...any) func(...uint) any {
	panic("flooring operations are not yet implemented")
}

func Ceiling(operands ...any) func(...uint) any {
	panic("ceiling operations are not yet implemented")
}

func Bound(minimum, maximum any, clamp bool, operands ...any) func(...uint) any {
	panic("bounding operations are not yet implemented")
}

func BaseToBase(base uint16, operands ...any) func(...uint) any {
	panic("base-to-base operations are not yet implemented")
}

// ParseInto takes in any operand and evaluates it into the provided type.  When TOut is a string, this will perform
// immediate base-to-base conversion.  If provided a string formula operand, this will execute it and yield the desired
// output.  You may optionally set the base for calculation to be performed in. If omitted (or if the output type is a
// num.Primitive), it will calculate in base₁₀
//
// To set the base, please provide a value from 0-256 to the first 'config' parameter.  To set the precision, please
// provide a value to the second 'config' parameter.
//
// NOTE: If you provide a base of 0 or 1, this will output the result as a string formula.  If provided a
// base₁, it will output with all identifiers fully expanded to their placeholder values - otherwise,
// named operands will be printed by their identifier.
//
// I.E. "42 + π" (base₀) or "42 + ~3.1415..." (base₁)
func ParseInto[TOut num.Advanced](operand any, config ...uint) TOut {
	panic("parsing operations are not yet implemented")

	// If TOut is a primitive, that's obvious
	// If TOut is a string, it should do base-to-base conversion
}
