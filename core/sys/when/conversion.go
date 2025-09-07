package when

import (
	"core/sys/num"
	"fmt"
	"time"
)

// DurationToHertz converts a time.Duration into Hertz (f = 1 / t)
func DurationToHertz[T num.Advanced](d time.Duration) T {
	var zero T
	switch typed := any(zero).(type) {
	case num.Natural, num.Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return durationToHertzAdvanced[T](d)
	case int:
		return any(durationToHertzPrimitive[int](d)).(T)
	case int8:
		return any(durationToHertzPrimitive[int8](d)).(T)
	case int16:
		return any(durationToHertzPrimitive[int16](d)).(T)
	case int32:
		return any(durationToHertzPrimitive[int32](d)).(T)
	case int64:
		return any(durationToHertzPrimitive[int64](d)).(T)
	case uint:
		return any(durationToHertzPrimitive[uint](d)).(T)
	case uint8:
		return any(durationToHertzPrimitive[uint8](d)).(T)
	case uint16:
		return any(durationToHertzPrimitive[uint16](d)).(T)
	case uint32:
		return any(durationToHertzPrimitive[uint32](d)).(T)
	case uint64:
		return any(durationToHertzPrimitive[uint64](d)).(T)
	case uintptr:
		return any(durationToHertzPrimitive[uintptr](d)).(T)
	case float32:
		return any(durationToHertzPrimitive[float32](d)).(T)
	case float64:
		return any(durationToHertzPrimitive[float64](d)).(T)
	default:
		panic(fmt.Errorf("unknown type %T", typed))
	}
}

func durationToHertzAdvanced[T num.Advanced](d time.Duration) T {
	panic("not implemented yet")
}

func durationToHertzPrimitive[T num.Primitive](d time.Duration) T {
	if d < 0 {
		d = 0
	}
	s := T(d) / 1e9
	hz := 1 / s
	return hz
}

// HertzToDuration converts a Hertz value to a time.Duration.
func HertzToDuration[T num.Primitive](hz T) time.Duration {
	switch typed := any(hz).(type) {
	case num.Natural, num.Realized, complex64, complex128:
		// CRITICAL: This does NOT type assert the string advanced types!!!
		// ALWAYS pass through the generic type for tiny types - never type assert and then mutate
		return hertzToDurationAdvanced(hz)
	case int:
		return hertzToDurationPrimitive(typed)
	case int8:
		return hertzToDurationPrimitive(typed)
	case int16:
		return hertzToDurationPrimitive(typed)
	case int32:
		return hertzToDurationPrimitive(typed)
	case int64:
		return hertzToDurationPrimitive(typed)
	case uint:
		return hertzToDurationPrimitive(typed)
	case uint8:
		return hertzToDurationPrimitive(typed)
	case uint16:
		return hertzToDurationPrimitive(typed)
	case uint32:
		return hertzToDurationPrimitive(typed)
	case uint64:
		return hertzToDurationPrimitive(typed)
	case uintptr:
		return hertzToDurationPrimitive(typed)
	case float32:
		return hertzToDurationPrimitive(typed)
	case float64:
		return hertzToDurationPrimitive(typed)
	default:
		panic(fmt.Errorf("unknown type %T", typed))
	}
}

func hertzToDurationAdvanced[T num.Advanced](hz T) time.Duration {
	panic("not implemented yet")
}

func hertzToDurationPrimitive[T num.Primitive](hz T) time.Duration {
	if hz <= 0 {
		// No division by zero
		switch typed := any(hz).(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, uintptr:
			hz = 1
		case float32, float64:
			typed = 1e-100 // math.SmallestNonzeroFloat64 🡨 NOTE: Raspberry Pi doesn't handle this constant well
			hz = typed
		default:
			panic("invalid type")
		}
	}
	s := 1 / hz
	ns := s * 1e9
	return time.Duration(ns)
}

// AbsDuration returns the absolute value of the provided duration.
func AbsDuration(d time.Duration) time.Duration {
	if d < 0 {
		d = -d
	}
	return d
}
