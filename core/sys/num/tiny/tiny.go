package tiny

import (
	"core/sys/num"
)

//func test[TA num.Advanced]() {
//	a := 6
//	c := Add[Real](a, 5)
//}

// Add takes in any number of Advanced objects and performs logical arithmetic upon them.  If either is not an Advanced type,
// this will panic - otherwise, the result will be provided in the requested Advanced type C.
func Add[TOut num.Advanced](operands ...any) TOut {
	var zero TOut
	return zero
}
