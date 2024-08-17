package JanOS

type Formula struct {
	Operator  string
	Operation func(source float64, variables ...float64) float64
}
