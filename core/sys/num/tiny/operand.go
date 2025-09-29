package tiny

import (
	"git.ignitelabs.net/janos/core/sys/given"
	"git.ignitelabs.net/janos/core/sys/given/format"
)

type Operand struct {
	Name   string
	Reveal func(...uint64) any
}

func Named(value any) Operand {
	return Operand{
		Name: given.Random[format.Tiny]().Name,
		Reveal: func(...uint64) any {
			// TODO: Parse the value and convert it to the requested base on the fly
			return value
		},
	}
}

var Pi = Operand{
	Name: "π",
	Reveal: func(...uint64) any {
		// TODO: Calculate this
		return "10#~3.14"
	},
}
