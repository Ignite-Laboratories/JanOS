package num

type RealMatrix struct {
	data []Real
}

func (m *RealMatrix) Align() {
	widest := uint(0)
	for _, r := range m.data {
		w := r.Width()
		if w > widest {
			widest = w
		}
	}

	for _, r := range m.data {

	}
}

func NewRealMatrix(operands ...any) RealMatrix {
	m := RealMatrix{
		data: make([]Real, len(operands)),
	}
	for i, operand := range operands {
		m.data[i] = NewReal(operand)
	}
	m.Align()
	return m
}
