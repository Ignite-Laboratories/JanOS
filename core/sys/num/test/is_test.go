package test

import (
	"core/sys/num"
	"math"
	"testing"
)

func Test_IsNaN(t *testing.T) {
	if !num.IsNaN(math.NaN()) {
		t.Errorf("a NaN was incorrectly detected as a number")
	}
	if num.IsNaN(i) {
		t.Errorf("a number was incorrectly detected as NaN")
	}
}

func Test_IsInf(t *testing.T) {
	neg := math.Inf(-1)
	nInf, nNeg := num.IsInf(neg)

	if nInf != true || nNeg != true {
		t.Errorf("IsInf(%v) = (%v, %v), want (%v, %v)", neg, nInf, nNeg, true, true)
	}

	pos := math.Inf(1)
	pInf, pNeg := num.IsInf(pos)

	if pInf != true || pNeg != false {
		t.Errorf("IsInf(%v) = (%v, %v), want (%v, %v)", pos, pInf, pNeg, true, false)
	}

	zero := math.Inf(0)
	zInf, zNeg := num.IsInf(zero)

	if zInf != true || zNeg != false {
		t.Errorf("IsInf(%v) = (%v, %v), want (%v, %v)", zero, zInf, zNeg, false, false)
	}

	nan := math.NaN()
	nInf, nNeg = num.IsInf(nan)

	if nInf != false || nNeg != false {
		t.Errorf("IsInf(%v) = (%v, %v), want (%v, %v)", nan, nInf, nNeg, false, false)
	}

	iInf, iNeg := num.IsInf(i)
	if iInf != false || iNeg != false {
		t.Errorf("a non-IEEE 754 type was incorrectly detected as 'infinity'")
	}
}

func Test_IsNumeric(t *testing.T) {
	if !num.IsNumeric(i, i8, i16, i32, i64, u, u8, u16, u32, u64, up, n, r, c64, c128) {
		t.Errorf("a numeric type was incorrectly detected as not a numeric")
	}

	if num.IsNumeric("", struct{}{}) {
		t.Errorf("a non-numeric type was incorrectly detected as a numeric")
	}
}

func Test_IsInteger(t *testing.T) {
	if !num.IsInteger(i, i8, i16, i32, i64, u, u8, u16, u32, u64, up, n) {
		t.Errorf("an integer type was incorrectly detected as not an integer")
	}

	if num.IsInteger(f32, f64, c64, c128, r) {
		t.Errorf("a non-integer type was incorrectly detected as an integer")
	}
}

func Test_IsComplex(t *testing.T) {
	if !num.IsComplex(c64, c128) {
		t.Errorf("a complex type was incorrectly detected as not a complex")
	}

	if num.IsFloat(i, i8, i16, i32, i64, u, u8, u16, u32, u64, up, n, r) {
		t.Errorf("a non-complex type was incorrectly detected as a complex")
	}
}

func Test_IsFloat(t *testing.T) {
	if !num.IsFloat(f32, f64, c64, c128, r) {
		t.Errorf("a float type was incorrectly detected as not a float")
	}

	if num.IsFloat(i, i8, i16, i32, i64, u, u8, u16, u32, u64, up, n) {
		t.Errorf("a non-float type was incorrectly detected as a float")
	}
}

func Test_IsSigned(t *testing.T) {
	if !num.IsSigned(i, i8, i16, i32, i64, f32, f64, c64, c128, r) {
		t.Errorf("a signed type was incorrectly detected as unsigned")
	}

	if num.IsSigned(u, u8, u16, u32, u64, up, n) {
		t.Errorf("an unsigned type was incorrectly detected as signed")
	}
}
