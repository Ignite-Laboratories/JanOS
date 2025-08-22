package num

import (
	"github.com/ignite-laboratories/core/enum/direction/orthogonal"
	"github.com/ignite-laboratories/core/sys/pad"
	"math"
	"regexp"
)

var decimalPattern = regexp.MustCompile(`^([+-]?)(\d+)(\.(\d+))?$`)

// Smallest returns the smaller of the two provided operands.
//
// NOTE: This takes in mixed types and casts the smaller operand to TOut - which can over and underflow.
func Smallest[TOut Primitive](a, b any) TOut {
	result := Compare(a, b)
	if result < 0 {
		return Cast[TOut](a)
	}
	return Cast[TOut](b)
}

// Largest returns the larger of the two provided operands.
//
// NOTE: This takes in mixed types and casts the smaller operand to TOut - which can over and underflow.
func Largest[TOut Primitive](a, b any) TOut {
	result := Compare(a, b)
	if result > 0 {
		return Cast[TOut](a)
	}
	return Cast[TOut](b)
}

// Compare performs a base-10 string comparison of whether 𝑎 is less than (-1), equal to (0), or greater than (1) 𝑏.
//
// NOTE: If working with floating point types, the following rules are applied:
//
//	NaN: Return whichever IS a number, otherwise panic
//	Inf: Is represented as "± MaxValue[uint64]() * 10" for comparison purposes
func Compare(a, b any) int {
	if !IsPrimitive(a, b) {
		panic("cannot compare non-primitive types")
	}

	aStr := ToString(a)
	bStr := ToString(b)

	if aStr == strNaN || bStr == strNaN {
		if aStr != strNaN {
			return 1
		} else if bStr != strNaN {
			return -1
		}
		panic("cannot compare " + strNaN)
	}

	if len(aStr) > 0 && aStr[1:] == strInf {
		if aStr[0] == '-' {
			aStr = "-" + ToString(MaxValue[uint64]()) + "0"
		}
		aStr = ToString(MaxValue[uint64]()) + "0"
	}

	aParts := decimalPattern.FindStringSubmatch(aStr)
	bParts := decimalPattern.FindStringSubmatch(bStr)
	if aParts == nil || bParts == nil {
		panic("unknown input type for comparison")
	}

	sign := 1
	whole := 2
	fractional := 4

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

	wholeSize := uint(math.Max(float64(len(aParts[whole])), float64(len(bParts[whole]))))
	fractionalSize := uint(math.Max(float64(len(aParts[fractional])), float64(len(bParts[fractional]))))

	aWhole := pad.String[rune, orthogonal.Left](wholeSize, aParts[whole], "0")
	bWhole := pad.String[rune, orthogonal.Left](wholeSize, bParts[whole], "0")

	aFractional := pad.String[rune, orthogonal.Right](fractionalSize, aParts[fractional], "0")
	bFractional := pad.String[rune, orthogonal.Right](fractionalSize, bParts[fractional], "0")

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
