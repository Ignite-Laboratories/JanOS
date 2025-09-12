package main

import (
	"fmt"
	"math/big"
	"reflect"

	"tiny5.0/atlas"
	"tiny5.0/test"
)

func main() {
	pass, _, providers := test.BuildTestCases()

	ops := filterOperands(pass...)
	ops = append(ops, filterOperands(providers...)...)

	fmt.Println("Hello, world!")
}

func filterOperands(operands ...any) []any {
	var filter func(any) any
	filter = func(op any) any {
		switch raw := op.(type) {

		// 0 - "Pass" branch
		case string,
			int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64,
			test.Realized, test.Natural, test.Measurement:
			return op

		// 1 - "Fail" branches
		case big.Int, big.Float:
			panic("big types should be pointers for normal operation")
		case big.Rat, *big.Rat:
			panic("big.Rat should use vector types")

		// 2 - "Recurse" branches
		case *big.Int:
			return filter(raw.Text(10))
		case *big.Float:
			return filter(raw.Text('f', int(atlas.Precision)))
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
