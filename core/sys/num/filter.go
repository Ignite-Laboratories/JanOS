package num

import (
	"core/std"
	"core/sys/atlas"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strings"
)

// filterOperands filters the provided operands into processable types.  If provided a type that does not
// satisfy the following requirements, this will panic -
//
//	0 - int, int8, int16, int32, int64 - Calls ToString
//	1 - uint, uint8, uint16, uint32, uint64, uintptr - Calls ToString
//	2 - float32, float64 - Calls ToString and panics on Inf or NaN
//	3 - big.Int, big.Float - Calls Text
//	4 - num.Realized, num.Realization, num.Natural, num.byte, num.Measurement - Passes through
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
// or invoking the operand until reaching the underlying value.  If you close over this function call, you dynamically
// encode in that functionality 'on the fly' to your code =)
func filterOperands(base uint16, operands ...any) []any {
	var filter func(any) any
	filter = func(op any) any {
		switch raw := op.(type) {

		// 0 - "Pass" branch
		case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
			return ToString(op)
		case float32:
			if math.IsInf(float64(raw), 0) {
				panic(fmt.Sprintf("cannot process an Inf valued %T", raw))
			}
			if math.IsNaN(float64(raw)) {
				panic(fmt.Sprintf("cannot process an NaN valued %T", raw))
			}
			return ToString(raw)
		case float64:
			if math.IsInf(raw, 0) {
				panic(fmt.Sprintf("cannot process an Inf valued %T", raw))
			}
			if math.IsNaN(raw) {
				panic(fmt.Sprintf("cannot process an NaN valued %T", raw))
			}
			return ToString(raw)
		case Realized, Natural, Measurement, Realization:
			return raw
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
			return ToString(raw)
		case *big.Float:
			return ToString(raw)
		case std.Neuron:
			return raw.Reveal()
		case *std.Neuron:
			return (*raw).Reveal()
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
