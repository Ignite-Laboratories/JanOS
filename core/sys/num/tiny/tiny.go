package tiny

import (
	"core/sys/atlas"
	"core/sys/num"
	"fmt"
	"reflect"
)

func sanityCheck(operands ...any) bool {
	if len(operands) == 0 {
		return
	}

	if !atlas.EnableRealCoercion {
		first := reflect.TypeOf(operands[0])

		for _, v := range operands[1:] {
			if reflect.TypeOf(v) != first {
				panic("with 'atlas.EnableReal = false', all operands must be of the same type")
			}
		}

		if !num.IsPrimitive(operands) {
			panic("with 'atlas.EnableReal = false', all operands must be of num.Primitive type")
		}
		return false
	}
	return true
}

//func test[TA num.Advanced]() {
//	a := 6
//	c := Add[Realized](a, 5)
//}

// NOTE NOTE NOTE NOTE!!!!!!
// ALEX!  Use num.Advanced for your output type!  This allows you to output a Realized, Realization, Measurement, or Primitive =)

// Add takes in any number of Advanced objects and performs logical arithmetic upon them.  If either is not an Advanced type,
// this will panic - otherwise, the result will be provided in the requested Advanced type C.
func Add[TOut num.Advanced](a any, b any, precision ...uint) TOut {
	realEnabled := sanityCheck(a, b)
	prec := atlas.Precision
	if precision != nil && len(precision) > 0 {
		prec = precision[0]
	}
	fmt.Println(prec)

	if realEnabled {
		// If here, bump all operands to num.Realized

		var zero TOut
		return zero
	}

	// If here, the types are all guaranteed to be like-typed primitives
	var zero TOut
	return zero
}
