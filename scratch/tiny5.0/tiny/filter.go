package tiny

import (
	"core/sys/atlas"
	"core/sys/num"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strings"
)

// FilterOperands filters the provided operands into processable types.  If provided a type that does not
// satisfy the following requirements, this will panic -
//
//	0 - int, int8, int16, int32, int64 - Calls num.ToString
//	1 - uint, uint8, uint16, uint32, uint64, uintptr - Calls num.ToString
//	2 - float32, float64 - Panics on Inf or NaN, then calls num.ToString
//	3 - big.Int, big.Float - Calls big.Text
//	4 - num.Realized, num.Realization, num.Measurement - Passes through
//	5 - string - Passes through
//	6 - []byte - Converts to a Natural as 'digits'
//
//	6 - pointers to any of the above types
//
//	7 - func() any or func[T any]() T
//	8 - func(uint) any or func[T any](uint) T
//	9 - func(*uint) any or func[T any](*uint) T
//	10 - func(...uint) any or func[T any](...uint) T
//	11 - func(...*uint) any or func[T any](...*uint) T
//
//	12 - Functions which return functions that satisfy the above functional requirements
//
// For function calls and pointer types, this will RESOLVE the underlying value they 'point' to by dereferencing
// or invoking the operand until reaching its result.  If you close over this function call, you dynamically
// encode in that functionality 'on the fly' to your code =)
//
// EXPECTED OUTPUT:
//
// This should either yield a parseable numeric string - see.PrintingNumbers - or, a
func FilterOperands(base uint16, operands ...any) []Realized {
	var filter func(any) Realized
	filter = func(op any) Realized {
		switch raw := op.(type) {

		// 0 - "Pass" branch
		case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
			return ParseRealized(num.ToString(op))
		case float32:
			if math.IsInf(float64(raw), 0) {
				panic(fmt.Sprintf("cannot process an Inf valued %T", raw))
			}
			if math.IsNaN(float64(raw)) {
				panic(fmt.Sprintf("cannot process an NaN valued %T", raw))
			}
			return ParseRealized(num.ToString(raw))
		case float64:
			if math.IsInf(raw, 0) {
				panic(fmt.Sprintf("cannot process an Inf valued %T", raw))
			}
			if math.IsNaN(raw) {
				panic(fmt.Sprintf("cannot process an NaN valued %T", raw))
			}
			return ParseRealized(num.ToString(raw))
		case Realized:
			return raw
		case Realization:
			wqefrqr32q134r5
		case []byte:
			digits := make([]string, len(raw))
			for i, d := range raw {
				digits[i] = fmt.Sprintf("%02x", d)
			}
			if base > 16 {
				return ParseNatural(strings.Join(digits, " "), base)
			}
			return ParseNatural(strings.Join(digits, ""), base)

		// 1 - "Fail" branches
		case big.Int, big.Float:
			panic("big types should be pointers for normal operation")
		case big.Rat, *big.Rat:
			panic("big.Rat should use vector types")

		// 2 - "Recurse" branches
		case *string:
			return filter(raw)
		case *big.Int:
			return ParseRealized(num.ToString(raw))
		case *big.Float:
			return ParseRealized(num.ToString(raw))
		default:
			rv := reflect.ValueOf(raw)
			if !rv.IsValid() {
				panic(fmt.Errorf("invalid type %T", raw))
			}
			for rv.Kind() == reflect.Pointer {
				if rv.IsNil() {
					panic(fmt.Errorf("got a nil input - %v", raw))
				}
				return filter(rv.Elem().Interface()) // Recurse!
			}

			switch rv.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return filter(rv.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				return filter(rv.Uint())
			case reflect.Float32, reflect.Float64:
				return filter(rv.Float())
			case reflect.Func:
				if rv.IsNil() {
					panic(fmt.Errorf("got a nil input - %v", raw))
				}

				t := reflect.TypeOf(raw)
				parameterCount := t.NumIn()
				args := make([]reflect.Value, 0)
				if parameterCount > 1 {
					panic("too many inputs")
				} else if parameterCount == 1 {
					p := t.In(0)
					valid := false
					if p.Kind() == reflect.Uint {
						valid = true
						args = append(args, reflect.ValueOf(atlas.Precision))
					} else if p.Kind() == reflect.Pointer && p.Elem().Kind() == reflect.Uint {
						valid = true
						args = append(args, reflect.ValueOf(&atlas.Precision))
					} else if p.Kind() == reflect.Slice && p.Elem().Kind() == reflect.Uint {
						valid = true
						args = append(args, reflect.ValueOf([]uint{atlas.Precision}))
					} else if p.Kind() == reflect.Slice && p.Elem().Kind() == reflect.Pointer && p.Elem().Elem().Kind() == reflect.Uint {
						valid = true
						args = append(args, reflect.ValueOf([]*uint{&atlas.Precision}))
					}
					if !valid {
						panic("invalid input parameters")
					}
				}
				if t.NumOut() != 1 {
					panic("must have exactly one output")
				}

				var result reflect.Value

				if t.IsVariadic() {
					result = rv.CallSlice(args)[0]
				} else {
					result = rv.Call(args)[0]
				}
				if result.Kind() == reflect.Pointer {
					return filter(result.Elem().Interface())
				}
				return filter(result.Interface())
			default:
				panic(fmt.Errorf("unknown type %T", raw))
			}
		}
	}

	result := make([]any, len(operands))
	for i, op := range operands {
		result[i] = filter(op)
	}
	return result
}
