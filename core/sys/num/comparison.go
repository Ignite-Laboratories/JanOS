package num

import (
	"core/enum/direction/ordinal"
	"core/sys/pad"
	"core/sys/pad/scheme"
	"regexp"
)

var decimalPattern = regexp.MustCompile(`^([+-]?)(\d+)(\.(\d+))?$`)
var allZerosPattern = regexp.MustCompile(`^[+-]?(?:0+(?:\.0*)?|\.(?:0)+)$`)

// Smallest returns the smaller of the two provided operands.
//
// NOTE: This takes in mixed types and casts the smaller operand to TOut - which can over and underflow.
func Smallest[TOut Primitive](a, b any) TOut {
	result := Compare(a, b)
	if result < 0 {
		return TypeAssert[TOut](a)
	}
	return TypeAssert[TOut](b)
}

// Largest returns the larger of the two provided operands.
//
// NOTE: This takes in mixed types and casts the smaller operand to TOut - which can over and underflow.
func Largest[TOut Primitive](a, b any) TOut {
	result := Compare(a, b)
	if result > 0 {
		return TypeAssert[TOut](a)
	}
	return TypeAssert[TOut](b)
}

// Compare performs a base-10 string comparison of whether the value of 𝑎 is less than (-1), equal to (0), or greater than (1) the value of 𝑏.
//
// NOTE: If working with IEEE 754 floating point types, 'Inf' is treated as a finite value beyond the other operand's value
// and NaN panics when both operands are NaN (otherwise it returns whichever IS a number).
func Compare(a, b any) int {
	if !IsNumeric(a, b) {
		panic("cannot compare non Numeric-compatible types")
	}

	if IsComplex(a, b) {
		panic("cannot compare complex numbers")
	}

	if IsNaN(a) || IsNaN(b) {
		if !IsNaN(a) {
			return 1
		} else if !IsNaN(b) {
			return -1
		}
		panic("cannot compare " + strNaN)
	}

	aInf, aInfNeg := IsInf(a)
	bInf, bInfNeg := IsInf(b)

	if aInf && bInf {
		if aInfNeg != bInfNeg {
			if aInfNeg {
				return -1
			}
			return 1
		}
		return 0
	}

	if aInf {
		if aInfNeg {
			return -1
		}
		return 1
	}
	if bInf {
		if bInfNeg {
			return 1
		}
		return -1
	}

	aStr := ToString(a)
	bStr := ToString(b)

	aParts := decimalPattern.FindStringSubmatch(aStr)
	bParts := decimalPattern.FindStringSubmatch(bStr)
	if aParts == nil || bParts == nil {
		panic("unknown input type for comparison")
	}

	sign := 1
	whole := 2
	fractional := 4

	if allZerosPattern.MatchString(aStr) && len(aParts[sign]) > 0 {
		aParts[sign] = ""
	}

	if allZerosPattern.MatchString(bStr) && len(bParts[sign]) > 0 {
		bParts[sign] = ""
	}

	if len(aParts[sign]) > 0 && len(bParts[sign]) == 0 {
		return -1
	}
	if len(aParts[sign]) == 0 && len(bParts[sign]) > 0 {
		return 1
	}

	negative := false
	if len(aParts[sign]) > 0 && len(bParts[sign]) > 0 {
		negative = true
	}

	wholeSize := len(aParts[whole])
	if len(bParts[whole]) > wholeSize {
		wholeSize = len(bParts[whole])
	}
	fractionalSize := len(aParts[fractional])
	if len(bParts[fractional]) > fractionalSize {
		fractionalSize = len(bParts[fractional])
	}

	aWhole := pad.String[rune](scheme.Tile, ordinal.Negative, uint(wholeSize), aParts[whole], "0")
	bWhole := pad.String[rune](scheme.Tile, ordinal.Negative, uint(wholeSize), bParts[whole], "0")

	aFractional := pad.String[rune](scheme.Tile, ordinal.Positive, uint(fractionalSize), aParts[fractional], "0")
	bFractional := pad.String[rune](scheme.Tile, ordinal.Positive, uint(fractionalSize), bParts[fractional], "0")

	aCombined := aWhole + aFractional
	bCombined := bWhole + bFractional

	for i := 0; i < len(aCombined); i++ {
		digitA := aCombined[i]
		digitB := bCombined[i]

		if digitA == digitB {
			continue
		}

		if digitA < digitB {
			if negative {
				return 1
			}
			return -1
		}
		if negative {
			return -1
		}
		return 1
	}
	return 0
}
