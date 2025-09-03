package test

import (
	"core/sys/num"
	"math"
	"testing"
)

func Test_Minimums(t *testing.T) {
	_i := num.MinValue[int]()
	if _i != math.MinInt {
		t.Errorf("MinValue[int] = %v, want %v", _i, math.MinInt)
	}
	_i8 := num.MinValue[int8]()
	if _i8 != math.MinInt8 {
		t.Errorf("MinValue[int8] = %v, want %v", _i8, math.MinInt8)
	}
	_i16 := num.MinValue[int16]()
	if _i16 != math.MinInt16 {
		t.Errorf("MinValue[int16] = %v, want %v", _i16, math.MinInt16)
	}
	_i32 := num.MinValue[int32]()
	if _i32 != math.MinInt32 {
		t.Errorf("MinValue[int32] = %v, want %v", _i32, math.MinInt32)
	}
	_i64 := num.MinValue[int64]()
	if _i64 != math.MinInt64 {
		t.Errorf("MinValue[int64] = %v, want %v", _i64, math.MinInt64)
	}

	_u := num.MinValue[uint]()
	if _u != 0 {
		t.Errorf("MinValue[uint] = %v, want %v", _u, 0)
	}
	_u8 := num.MinValue[uint8]()
	if _u8 != 0 {
		t.Errorf("MinValue[uint8] = %v, want %v", _u8, 0)
	}
	_u16 := num.MinValue[uint16]()
	if _u16 != 0 {
		t.Errorf("MinValue[uint16] = %v, want %v", _u16, 0)
	}
	_u32 := num.MinValue[uint32]()
	if _u32 != 0 {
		t.Errorf("MinValue[uint32] = %v, want %v", _u32, 0)
	}
	_u64 := num.MinValue[uint64]()
	if _u64 != 0 {
		t.Errorf("MinValue[uint64] = %v, want %v", _u64, 0)
	}
	_up := num.MinValue[uintptr]()
	if _up != 0 {
		t.Errorf("MinValue[uintptr] = %v, want %v", _u64, 0)
	}

	_f32 := num.MinValue[float32]()
	if _f32 != 0 {
		t.Errorf("MinValue[float32] = %v, want %v", _f32, 0)
	}
	_f64 := num.MinValue[float64]()
	if _f64 != 0 {
		t.Errorf("MinValue[float64] = %v, want %v", _f64, 0)
	}
}

func Test_Maximums(t *testing.T) {
	_i := num.MaxValue[int]()
	if _i != math.MaxInt {
		t.Errorf("MaxValue[int] = %v, want %v", _i, math.MaxInt)
	}
	_i8 := num.MaxValue[int8]()
	if _i8 != math.MaxInt8 {
		t.Errorf("MaxValue[int8] = %v, want %v", _i8, math.MaxInt8)
	}
	_i16 := num.MaxValue[int16]()
	if _i16 != math.MaxInt16 {
		t.Errorf("MaxValue[int16] = %v, want %v", _i16, math.MaxInt16)
	}
	_i32 := num.MaxValue[int32]()
	if _i32 != math.MaxInt32 {
		t.Errorf("MaxValue[int32] = %v, want %v", _i32, math.MaxInt32)
	}
	_i64 := num.MaxValue[int64]()
	if _i64 != math.MaxInt64 {
		t.Errorf("MaxValue[int64] = %v, want %v", _i64, math.MaxInt64)
	}

	_u := num.MaxValue[uint]()
	if _u != math.MaxUint {
		t.Errorf("MaxValue[uint] = %v, want math.MaxUint", _u)
	}
	_u8 := num.MaxValue[uint8]()
	if _u8 != math.MaxUint8 {
		t.Errorf("MaxValue[uint8] = %v, want %v", _u8, math.MaxUint8)
	}
	_u16 := num.MaxValue[uint16]()
	if _u16 != math.MaxUint16 {
		t.Errorf("MaxValue[uint16] = %v, want %v", _u16, math.MaxUint16)
	}
	_u32 := num.MaxValue[uint32]()
	if _u32 != math.MaxUint32 {
		t.Errorf("MaxValue[uint32] = %v, want %v", _u32, math.MaxUint32)
	}
	_u64 := num.MaxValue[uint64]()
	if _u64 != math.MaxUint64 {
		t.Errorf("MaxValue[uint64] = %v, want math.MaxUint64", _u64)
	}
	_up := num.MaxValue[uintptr]()
	if _up != math.MaxUint {
		t.Errorf("MaxValue[uintptr] = %v, want math.MaxUint64", _up)
	}

	_f32 := num.MaxValue[float32]()
	if _f32 != 1 {
		t.Errorf("MaxValue[float32] = %v, want %v", _f32, 1)
	}
	_f64 := num.MaxValue[float64]()
	if _f64 != 1 {
		t.Errorf("MaxValue[float64] = %v, want %v", _f64, 1)
	}
}
