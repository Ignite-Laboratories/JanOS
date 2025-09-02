package tiny

type Placeholder string

func (p Placeholder) String() string {
	return string(p)
}

type Natural []Placeholder

func (n Natural) String() string {
	// TODO: Implement this
	return "0"
}

// Real asdf
type Real struct {
	Negative   bool
	Whole      Natural
	Fractional Natural
	Periodic   Natural
}

func (r Real) String() string {
	// TODO: Implement this
	return "0"
}

//func test[TA num.Advanced]() {
//	a := 6
//	c := Add[Real](a, 5)
//}

// Add takes in any number of Advanced objects and performs logical arithmetic upon them.  If either is not an Advanced type,
// this will panic - otherwise, the result will be provided in the requested Advanced type C.
//func Add[C num.Advanced](operands ...any) C {
//
//}
