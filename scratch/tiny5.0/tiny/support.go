package tiny

import "fmt"

/*
Error messages
*/

// panicIfInvalidBase will return base₁₀ if no input is provided, or panic if it's not in the closed set [base₂, base₂₅₆]
func panicIfInvalidBase(current uint16, assignment ...uint16) uint16 {
	if len(assignment) == 0 {
		if assignment[0] < 2 || assignment[0] > 256 {
			panic(fmt.Errorf("invalid assignment '%d' - must be between 2 and 256", assignment[0]))
		}
		current = assignment[0]
	}
	return current
}
