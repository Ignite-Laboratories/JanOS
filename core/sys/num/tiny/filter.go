package tiny

import (
	"fmt"
	"math"
	"math/big"
	"reflect"

	"git.ignitelabs.net/janos/core/sys/atlas"
	"git.ignitelabs.net/janos/core/sys/num"
)

// FilterOperands filters the provided operands into processable types.  If provided a type that does not
// satisfy the following requirements, this will panic -
//
//		0 - int, int8, int16, int32, int64 - Calls num.ToString
//		1 - uint, uint8, uint16, uint32, uint64, uintptr - Calls num.ToString
//		2 - float32, float64 - Panics on Inf or NaN, then calls num.ToString
//		3 - big.Int, big.Float - Calls big.Text
//		5 - string - Passes through
//		6 - []byte - Passes through - This is treated as a natural number with each byte a placeholder
//	 	7 - num.Bounds - Passes through
//
//		8 - pointers to any of the above types are dereferenced and treated as above
//
//		Function call types:
//
//		9 - func() any or func[T any]() T
//		10 - func(base uint) any or func[T any](base uint) T
//		11 - func(base *uint) any or func[T any](base *uint) T
//		12 - func(base ...uint) any or func[T any](base ...uint) T
//		13 - func(base ...*uint) any or func[T any](base ...*uint) T
//
// For function calls and pointer types, this will RESOLVE the underlying value they 'point' to by dereferencing
// or invoking the operand until reaching its result.  If you close over this function call, you dynamically
// encode in that functionality 'on the fly' =)
func FilterOperands(base uint16, operands ...any) []any {
	var filter func(any) any
	filter = func(op any) any {
		switch raw := op.(type) {

		// 0 - "Pass" branch
		case num.Bounds:
			return raw
		case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
			return num.ToString(raw)
		case float32:
			if math.IsInf(float64(raw), 0) {
				panic(fmt.Sprintf("cannot process an Inf valued %T", raw))
			}
			if math.IsNaN(float64(raw)) {
				panic(fmt.Sprintf("cannot process an NaN valued %T", raw))
			}
			return num.ToString(raw)
		case float64:
			if math.IsInf(raw, 0) {
				panic(fmt.Sprintf("cannot process an Inf valued %T", raw))
			}
			if math.IsNaN(raw) {
				panic(fmt.Sprintf("cannot process an NaN valued %T", raw))
			}
			return num.ToString(raw)
		case []byte:
			return raw
		// 1 - "Fail" branches
		case big.Int, big.Float:
			panic("big types should be pointers for normal operation")
		case big.Rat, *big.Rat:
			panic("big.Rat should use vector types")

		// 2 - "Recurse" branches
		case *string:
			return filter(raw)
		case *big.Int:
			// TODO: big doesn't cover all of tiny's bases, so we still need to do a base conversion from big's output
			return raw.Text(10)
		case *big.Float:
			// TODO: big doesn't cover all of tiny's bases, so we still need to do a base conversion from big's output
			return raw.Text(10, int(atlas.Precision))
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
