package Formula

import (
	"github.com/ignite-laboratories/JanOS/util"
)

var Additive = util.Formula{
	Operator:  "+",
	Operation: additive,
}

func additive(source float64, variables ...float64) float64 {
	val := source

	for _, variable := range variables {
		val += variable
	}

	return val
}

var Subtractive = util.Formula{
	Operator:  "-",
	Operation: subtractive,
}

func subtractive(source float64, variables ...float64) float64 {
	val := source

	for _, variable := range variables {
		val -= variable
	}

	return val
}

var Multiplicative = util.Formula{
	Operator:  "*",
	Operation: multiplicative,
}

func multiplicative(source float64, variables ...float64) float64 {
	val := source

	for _, variable := range variables {
		val *= variable
	}

	return val
}

var Divisive = util.Formula{
	Operator:  "/",
	Operation: divisive,
}

func divisive(source float64, variables ...float64) float64 {
	val := source

	for _, variable := range variables {
		val /= variable
	}

	return val
}
