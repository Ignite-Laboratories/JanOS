package Formula

import (
	"github.com/ignite-laboratories/JanOS"
)

var Additive = JanOS.Formula{
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

var Subtractive = JanOS.Formula{
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

var Multiplicative = JanOS.Formula{
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

var Divisive = JanOS.Formula{
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
