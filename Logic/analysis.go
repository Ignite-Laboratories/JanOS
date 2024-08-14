package Logic

import "gonum.org/v1/gonum/mat"

type Analyzer interface {
	Analyze(dense mat.Dense)
}
