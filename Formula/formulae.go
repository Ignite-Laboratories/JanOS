package Formula

import (
	"github.com/ignite-laboratories/JanOS"
	"time"
)

var Additive = JanOS.Formula{
	Operator:  "+",
	Operation: additive,
}

func additive(instant time.Time, sourceSignal *JanOS.Signal, otherSignals ...*JanOS.Signal) float64 {
	val := sourceSignal.GetValue(instant).Value

	for _, otherSignal := range otherSignals {
		val += otherSignal.GetValue(instant).Value
	}

	return val
}

var Subtractive = JanOS.Formula{
	Operator:  "-",
	Operation: subtractive,
}

func subtractive(instant time.Time, sourceSignal *JanOS.Signal, otherSignals ...*JanOS.Signal) float64 {
	val := sourceSignal.GetValue(instant).Value

	for _, otherSignal := range otherSignals {
		val -= otherSignal.GetValue(instant).Value
	}

	return val
}

var Multiplicative = JanOS.Formula{
	Operator:  "*",
	Operation: multiplicative,
}

func multiplicative(instant time.Time, sourceSignal *JanOS.Signal, otherSignals ...*JanOS.Signal) float64 {
	val := sourceSignal.GetValue(instant).Value

	for _, otherSignal := range otherSignals {
		val *= otherSignal.GetValue(instant).Value
	}

	return val
}

var Divisive = JanOS.Formula{
	Operator:  "/",
	Operation: divisive,
}

func divisive(instant time.Time, sourceSignal *JanOS.Signal, otherSignals ...*JanOS.Signal) float64 {
	val := sourceSignal.GetValue(instant).Value

	for _, otherSignal := range otherSignals {
		val /= otherSignal.GetValue(instant).Value
	}

	return val
}
