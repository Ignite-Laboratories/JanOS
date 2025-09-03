package num

import (
	"fmt"
	"math"
	"reflect"
)

// MaxValue returns an instance of T set to the maximum value of the provided Primitive type.
//
// NOTE: Floating point and complex types treat their implicit "boundaries" as [0.0, 1.0] (including the imaginary part)
func MaxValue[T Primitive]() T {
	return boundaryValue[T](false)
}

// MinValue returns an instance of T set to the minimum value of the provided Primitive type.
//
// NOTE: Floating point and complex types treat their implicit "boundaries" as [0.0, 1.0] (including the imaginary part)
func MinValue[T Primitive]() T {
	return boundaryValue[T](true)
}

func boundaryValue[T Primitive](min bool) T {
	var out uint64
	var zero T
	switch any(zero).(type) {
	case int8:
		if min {
			v := math.MinInt8
			return T(v)
		}
		out = math.MaxInt8
	case uint8:
		if min {
			return 0
		}
		out = math.MaxUint8
	case int16:
		if min {
			v := math.MinInt16
			return T(v)
		}
		out = math.MaxInt16
	case uint16:
		if min {
			return 0
		}
		out = math.MaxUint16
	case int32:
		if min {
			v := math.MinInt32
			return T(v)
		}
		out = math.MaxInt32
	case uint32:
		if min {
			return 0
		}
		out = math.MaxUint32
	case int64:
		if min {
			v := math.MinInt64
			return T(v)
		}
		out = math.MaxInt64
	case uint64:
		if min {
			return 0
		}
		out = math.MaxUint64
	case int:
		if min {
			v := math.MinInt
			return T(v)
		}
		out = math.MaxInt
	case uint:
		if min {
			return 0
		}
		out = math.MaxUint
	case uintptr:
		if min {
			return 0
		}
		out = math.MaxUint64
	case float32, float64, complex64, complex128:
		if min {
			return T(0)
		}
		return T(1)
	default:
		switch reflect.TypeOf(T(0)).Kind() {
		case reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
			if min {
				return T(0)
			}
			return T(1)
		case reflect.Int8:
			if min {
				v := math.MinInt8
				return T(v)
			}
			out = math.MaxInt8
		case reflect.Int16:
			if min {
				v := math.MinInt16
				return T(v)
			}
			out = math.MaxInt16
		case reflect.Int32:
			if min {
				v := math.MinInt32
				return T(v)
			}
			out = math.MaxInt32
		case reflect.Int64:
			if min {
				v := math.MinInt64
				return T(v)
			}
			out = math.MaxInt64
		case reflect.Int:
			if min {
				v := math.MinInt
				return T(v)
			}
			out = math.MaxInt
		case reflect.Uint8:
			if min {
				return 0
			}
			out = math.MaxUint8
		case reflect.Uint16:
			if min {
				return 0
			}
			out = math.MaxUint16
		case reflect.Uint32:
			if min {
				return 0
			}
			out = math.MaxUint32
		case reflect.Uint64:
			if min {
				return 0
			}
			out = math.MaxUint64
		case reflect.Uint:
			if min {
				return 0
			}
			out = math.MaxUint
		case reflect.Uintptr:
			if min {
				return 0
			}
			out = math.MaxUint64
		default:
			panic(fmt.Errorf("unknown type %T", zero))
		}
	}
	return T(out)
}
