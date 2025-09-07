package internal

// Irrational is a container for the internal irrationality system.
type Irrational struct {
	irrational func(uint)
}

func NewIrrational(fn func(uint)) Irrational {
	return Irrational{
		irrational: fn,
	}
}

// Is indicates if the num.Realized value is an irrational number dynamically rendered at runtime.
func (i *Irrational) Is() bool {
	return i.irrational != nil
}
