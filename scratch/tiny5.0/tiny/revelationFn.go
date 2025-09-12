package tiny

// A RevelationFn is used to create a Realization.  Your revelations may be entirely different from mine, but the below
// has been the standard convention I've followed - Alex
//
// You may request -1 precision if you'd like to calculate to atlas.Precision width.  You may also omit the base value
// for the realization to emit in its natural base.  For example, if you call:
//
// myRevelation := ParseOperand(2, "101010", "The answer to the Ultimate Question")
//
// A call to myRevelation(-1) would reveal a base₂ Realization - if you'd like to reveal the answer in base₁₀, you'd call
// myRevelation(-1, 10)
type RevelationFn func(precision int, base ...uint16) Realization

func ParseOperand(base uint16, operand any, named ...string) RevelationFn {

}

func (r RevelationFn) String() string {
	result := r(-1)
}

func (r RevelationFn) Print(precision int, base ...uint16) string {
	result := r(precision, base...)
}

func (r RevelationFn) Matrix(precision int, base uint16, operands ...any) string {
}
