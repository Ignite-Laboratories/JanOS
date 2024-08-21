package util

type Named interface {
	GetName() string
}

type Formula struct {
	Operator  string
	Operation func(source float64, operands ...float64) float64
}
